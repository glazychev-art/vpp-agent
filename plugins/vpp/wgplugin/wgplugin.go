
//go:generate descriptor-adapter --descriptor-name Device --value-type *vpp_wg.Device --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg" --output-dir "descriptor"
//go:generate descriptor-adapter --descriptor-name Peer --value-type *vpp_wg.Peer --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg" --output-dir "descriptor"

package wgplugin

import (
	"github.com/pkg/errors"
	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/infra"
	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/descriptor"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/vppcalls"

	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/vppcalls/vpp2001"
)

type WgPlugin struct {
	Deps
	// handler
	WgHandler vppcalls.WgVppAPI

	deviceDescriptor *descriptor.WgDeviceDescriptor
	peerDescriptor *descriptor.WgPeerDescriptor
}

type Deps struct {
	infra.PluginDeps
	KVScheduler kvs.KVScheduler
	VPP         govppmux.API
	IfPlugin    ifplugin.API
	StatusCheck statuscheck.PluginStatusWriter // optional
}

func (p *WgPlugin) Init() (err error) {
	// init Wg handler
	p.WgHandler = vppcalls.CompatibleWgVppHandler(p.VPP, p.IfPlugin.GetInterfaceIndex(), p.Log)
	if p.WgHandler == nil {
		return errors.New("WgHandler is not available")
	}

	p.deviceDescriptor = descriptor.NewWgDeviceDescriptor(p.WgHandler, p.Log)
	deviceDescriptor := adapter.NewDeviceDescriptor(p.deviceDescriptor.GetDescriptor())
	err = p.KVScheduler.RegisterKVDescriptor(deviceDescriptor)
	if err != nil {
		return err
	}

	p.peerDescriptor = descriptor.NewWgPeerDescriptor(p.WgHandler, p.Log)
	peerDescriptor := adapter.NewPeerDescriptor(p.peerDescriptor.GetDescriptor())
	err = p.KVScheduler.RegisterKVDescriptor(peerDescriptor)
	if err != nil {
		return err
	}

	return nil
}

// AfterInit registers plugin with StatusCheck.
func (p *WgPlugin) AfterInit() error {
	if p.StatusCheck != nil {
		p.StatusCheck.Register(p.PluginName, nil)
	}
	return nil
}
