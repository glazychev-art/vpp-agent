package descriptor

import (
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/pkg/models"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"

	//kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/wgplugin/vppcalls"
	wg "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg"
)

const (
	// DeviceDescriptorName is the name of the descriptor for VPP wg device.
	DeviceDescriptorName = "vpp-wg-device"
)

// WgDeviceDescriptor teaches KVScheduler how to configure VPP Wg device.
type WgDeviceDescriptor struct {
	log          logging.Logger
	wgHandler vppcalls.WgVppAPI
}

// NewWgDeviceDescriptor creates a new instance of the Wg Device descriptor.
func NewWgDeviceDescriptor(wgHandler vppcalls.WgVppAPI, log logging.PluginLogger) *WgDeviceDescriptor {
	return &WgDeviceDescriptor{
		wgHandler:    wgHandler,
		log:          log.NewLogger("wg-device-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (d *WgDeviceDescriptor) GetDescriptor() *adapter.DeviceDescriptor {
	return &adapter.DeviceDescriptor{
		Name:            DeviceDescriptorName,
		NBKeyPrefix:     wg.ModelDevice.KeyPrefix(),
		ValueTypeName:   wg.ModelDevice.ProtoName(),
		KeySelector:     wg.ModelDevice.IsKeyValid,
		KeyLabel:        wg.ModelDevice.StripKeyPrefix,
		ValueComparator: d.EquivalentWgDevices,
		Create:          d.Create,
		Delete:          d.Delete,
		Retrieve:        d.Retrieve,
	}
}

func (d *WgDeviceDescriptor) EquivalentWgDevices(key string, oldDevice, newDevice *wg.Device) bool {
	// compare base fields
	return oldDevice.PrivateKey == newDevice.PrivateKey &&
		oldDevice.Port == newDevice.Port
}

// Create adds a new security association pair.
func (d *WgDeviceDescriptor) Create(key string, device *wg.Device) (metadata interface{}, err error) {
	// add security association
	err = d.wgHandler.SetDevice(device)
	if err != nil {
		d.log.Error(err)
	}

	return nil, err
}

// Delete removes VPP wg device.
func (d *WgDeviceDescriptor) Delete(key string, device *wg.Device, metadata interface{}) error {
	err := d.wgHandler.RemoveDevice()
	if err != nil {
		d.log.Error(err)
	}
	return err
}

// Retrieve returns vpp wg device
func (d *WgDeviceDescriptor) Retrieve(correlate []adapter.DeviceKVWithMetadata) (dump []adapter.DeviceKVWithMetadata, err error) {
	device, err := d.wgHandler.DumpWgDevice()
	if err != nil {
		d.log.Error(err)
		return dump, err
	}
	if device.Port == 0 {
		return
	}
	dump = append(dump, adapter.DeviceKVWithMetadata{
		Key:      models.Key(device),
		Value:    device,
		Origin:   kvs.FromNB,
	})
	return dump, nil
}
