//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp2001

import (
	"fmt"
	vpp_wg "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/wg"
	wg "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg"
)

func (h *WgVppHandler) SetDevice(device *wg.Device) error {
	request := &vpp_wg.WgSetDevice{
		PrivateKey : []byte(device.PrivateKey),
		Port : uint16(device.Port),
	}
	// prepare reply
	reply := &vpp_wg.WgSetDeviceReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil;
}

func (h *WgVppHandler) RemoveDevice() error {
	// prepare request
	request := &vpp_wg.WgRemoveDevice{}
	// prepare reply
	reply := &vpp_wg.WgRemoveDeviceReply{}
	// send request and obtain reply

	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil;
}

func (h *WgVppHandler) SetPeer(peer *wg.Peer) error {
	request := &vpp_wg.WgSetPeer{
		PublicKey : []byte(peer.PublicKey),
		Port : uint16(peer.Port),
		PersistentKeepalive: uint16(peer.PersistentKeepalive),
	}

	ifaceMeta, found := h.ifIndexes.LookupByName(peer.TunInterface)
	if !found {
		return fmt.Errorf("failed to get interface metadata")
	}
	request.TunSwIfIndex = vpp_wg.InterfaceIndex(ifaceMeta.SwIfIndex)

	var err error
	request.Endpoint, err = IPToAddress(peer.Endpoint)
	if err != nil {
		return err
	}

	request.AllowedIP, err = IPToAddress(peer.AllowedIp)
	if err != nil {
		return err
	}

	// prepare reply
	reply := &vpp_wg.WgSetPeerReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil;
}

func (h *WgVppHandler) RemovePeer(peer *wg.Peer) error {
	// prepare request
	request := &vpp_wg.WgRemovePeer{
		PublicKey: []byte(peer.PublicKey),
	}
	// prepare reply
	reply := &vpp_wg.WgRemovePeerReply{}
	// send request and obtain reply

	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil;
}

