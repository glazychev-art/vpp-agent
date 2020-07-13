package descriptor

import (
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/pkg/models"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/vppcalls"
	wg "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg"
)

const (
	// PeerDescriptorName is the name of the descriptor for VPP wg peer.
	PeerDescriptorName = "vpp-wg-peer"
)

// WgPeerDescriptor teaches KVScheduler how to configure VPP Wg peer.
type WgPeerDescriptor struct {
	log          logging.Logger
	wgHandler vppcalls.WgVppAPI
}

// NewWgDeviceDescriptor creates a new instance of the Wg Device descriptor.
func NewWgPeerDescriptor(wgHandler vppcalls.WgVppAPI, log logging.PluginLogger) *WgPeerDescriptor {
	return &WgPeerDescriptor{
		wgHandler:    wgHandler,
		log:          log.NewLogger("wg-peer-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (d *WgPeerDescriptor) GetDescriptor() *adapter.PeerDescriptor {
	return &adapter.PeerDescriptor{
		Name:            PeerDescriptorName,
		NBKeyPrefix:     wg.ModelPeer.KeyPrefix(),
		ValueTypeName:   wg.ModelPeer.ProtoName(),
		KeySelector:     wg.ModelPeer.IsKeyValid,
		KeyLabel:        wg.ModelPeer.StripKeyPrefix,
		ValueComparator: d.EquivalentWgPeers,
		Create:          d.Create,
		Delete:          d.Delete,
		Retrieve:        d.Retrieve,
	}
}

func (d *WgPeerDescriptor) EquivalentWgPeers(key string, oldPeer, newPeer *wg.Peer) bool {
	// compare base fields
	return oldPeer.PublicKey == newPeer.PublicKey &&
		oldPeer.Port == newPeer.Port
}

// Create adds a new security association pair.
func (d *WgPeerDescriptor) Create(key string, peer *wg.Peer) (metadata interface{}, err error) {
	// add security association
	err = d.wgHandler.SetPeer(peer)
	if err != nil {
		d.log.Error(err)
	}

	return nil, err
}

// Delete removes VPP wg device.
func (d *WgPeerDescriptor) Delete(key string, peer *wg.Peer, metadata interface{}) error {
	err := d.wgHandler.RemovePeer(peer)
	if err != nil {
		d.log.Error(err)
	}
	return err
}

// Retrieve returns all wg peers.
func (d *WgPeerDescriptor) Retrieve(correlate []adapter.PeerKVWithMetadata) (dump []adapter.PeerKVWithMetadata, err error) {
	peers, err := d.wgHandler.DumpWgPeers()
	if err != nil {
		d.log.Error(err)
		return dump, err
	}
	for _, peer := range peers {
		dump = append(dump, adapter.PeerKVWithMetadata{
			Key:      models.Key(peer),
			Value:    peer,
			Origin:   kvs.FromNB,
		})
	}

	return dump, nil
}