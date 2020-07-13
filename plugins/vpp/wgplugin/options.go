
package wgplugin

import (
	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/logging"

	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	"go.ligato.io/vpp-agent/v3/plugins/kvscheduler"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
)

var DefaultPlugin = *NewPlugin()

// NewPlugin creates a new Plugin with the provided Options.
func NewPlugin(opts ...Option) *WgPlugin {
	p := &WgPlugin{}

	p.PluginName = "vpp-wg-plugin"
	p.StatusCheck = &statuscheck.DefaultPlugin
	p.KVScheduler = &kvscheduler.DefaultPlugin
	p.VPP = &govppmux.DefaultPlugin
	p.IfPlugin = &ifplugin.DefaultPlugin

	for _, o := range opts {
		o(p)
	}

	if p.Log == nil {
		p.Log = logging.ForPlugin(p.String())
	}

	return p
}

// Option is a function that can be used in NewPlugin to customize Plugin.
type Option func(*WgPlugin)

// UseDeps returns Option that can inject custom dependencies.
func UseDeps(f func(*Deps)) Option {
	return func(p *WgPlugin) {
		f(&p.Deps)
	}
}
