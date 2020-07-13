package vppcalls

import (
	govppapi "git.fd.io/govpp.git/api"
	"go.ligato.io/cn-infra/v2/logging"
	wg "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg"

	"go.ligato.io/vpp-agent/v3/plugins/vpp"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
)

type WgVppAPI interface {
	WgVppRead

	// Set device via binary API
	SetDevice(device *wg.Device) error
	// Remove device via binary API
	RemoveDevice() error
	// Set peer via binary API
	SetPeer(peer *wg.Peer) error
	// Remove peer via binary API
	RemovePeer(peer *wg.Peer) error
}

// WgVPPRead provides read methods for Wg
type WgVppRead interface {
	// DumpWgDevice returns a device state
	DumpWgDevice() (device *wg.Device, err error)
	// DumpWgPeers returns a peers states
	DumpWgPeers() (peerList []*wg.Peer, err error)
}

var Handler = vpp.RegisterHandler(vpp.HandlerDesc{
	Name:       "wg",
	HandlerAPI: (*WgVppAPI)(nil),
})

type NewHandlerFunc func(ch govppapi.Channel, ifDdx ifaceidx.IfaceMetadataIndex, log logging.Logger) WgVppAPI

func AddHandlerVersion(version vpp.Version, msgs []govppapi.Message, h NewHandlerFunc) {
	Handler.AddVersion(vpp.HandlerVersion{
		Version: version,
		Check: func(c vpp.Client) error {
			ch, err := c.NewAPIChannel()
			if err != nil {
				return err
			}
			return ch.CheckCompatiblity(msgs...)
		},
		NewHandler: func(c vpp.Client, a ...interface{}) vpp.HandlerAPI {
			ch, err := c.NewAPIChannel()
			if err != nil {
				return err
			}
			return h(ch, a[0].(ifaceidx.IfaceMetadataIndex), a[1].(logging.Logger))
		},
	})
}

func CompatibleWgVppHandler(c vpp.Client, ifIdx ifaceidx.IfaceMetadataIndex, log logging.Logger) WgVppAPI {
	if v := Handler.FindCompatibleVersion(c); v != nil {
		return v.NewHandler(c, ifIdx, log).(WgVppAPI)
	}
	return nil
}
