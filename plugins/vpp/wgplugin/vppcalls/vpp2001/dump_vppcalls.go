package vpp2001

import (
	"bytes"
	"net"

	vpp_wg "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/wg"
	wg "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg"
)

// DumpWgDevice implements wg handler.
func (h *WgVppHandler) DumpWgDevice() (device *wg.Device, err error) {
	req := &vpp_wg.WgDeviceDump{}
	deviceDetails := &vpp_wg.WgDeviceDetails{}

	err = h.callsChannel.SendRequest(req).ReceiveReply(deviceDetails)
	if err != nil {
		return nil, err
	}

	device = &wg.Device{
		PrivateKey: string(bytes.Trim(deviceDetails.PrivateKey, "\x00")),
		Port:       uint32(deviceDetails.Port),
	}
	return
}

// DumpWgPeers implements wg handler.
func (h *WgVppHandler) DumpWgPeers() (peerList []*wg.Peer, err error) {
	req := &vpp_wg.WgPeersDump{}
	requestCtx := h.callsChannel.SendMultiRequest(req)

	var vppPeerList []*vpp_wg.WgPeersDetails
	for {
		vppPeerDetails := &vpp_wg.WgPeersDetails{}
		stop, err := requestCtx.ReceiveReply(vppPeerDetails)
		if stop {
			break
		}
		if err != nil {
			return nil, err
		}
		vppPeerList = append(vppPeerList, vppPeerDetails)
	}

	for _, vppPeerDetails := range vppPeerList {
		addr := net.IP(vppPeerDetails.IP4Address[:])
		peerDetails := &wg.Peer{
			PublicKey: string(bytes.Trim(vppPeerDetails.PublicKey, "\x00")),
			Endpoint: addr.String(),
		}
		peerList = append(peerList, peerDetails)
	}

	return
}

