package vpp_wg

import (
	"go.ligato.io/vpp-agent/v3/pkg/models"
)

// ModuleName is the module name used for models.
const ModuleName = "vpp.wg"

var (
	ModelDevice = models.Register(&Device{}, models.Spec{
		Module:  ModuleName,
		Version: "v1",
		Type:    "device",
	})

	ModelPeer = models.Register(&Peer{}, models.Spec{
		Module:  ModuleName,
		Version: "v1",
		Type:    "peer",
	}, models.WithNameTemplate("{{.PublicKey}}"))
)
