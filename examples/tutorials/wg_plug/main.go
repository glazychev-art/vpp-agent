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

package main

import (
	"log"
	"net"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/wg"

	"git.fd.io/govpp.git/api"

	"go.ligato.io/vpp-agent/v3/plugins/govppmux"

	"go.ligato.io/cn-infra/v2/agent"
)

//go:generate binapi-generator --input-file=/home/art/xor/vpp/build-root/install-vpp-native/vpp/share/vpp/api/plugins/wg.api.json --output-dir=./bin_api22

func main() {
	// Create an instance of our plugin.
	p := new(WgPlugin)
	p.GoVPPMux = &govppmux.DefaultPlugin

	// Create new agent with our plugin instance.
	a := agent.NewAgent(agent.AllPlugins(p))

	// Run starts the agent with plugins, wait until shutdown
	// and then stops the agent and its plugins.
	if err := a.Start(); err != nil {
		log.Fatalln(err)
	}

	p.syncVppCall()

	if err := a.Stop(); err != nil {
		log.Fatalln(err)
	}
}

// HelloWorld represents our plugin.
type WgPlugin struct {
	vppChan api.Channel

	GoVPPMux govppmux.API
}

// String is used to identify the plugin by giving it name.
func (p *WgPlugin) String() string {
	return "HelloWorld"
}

// Init is executed on agent initialization.
func (p *WgPlugin) Init() (err error) {
	log.Println("====== Wg plugin init")

	if p.vppChan, err = p.GoVPPMux.NewAPIChannel(); err != nil {
		panic(err)
	}

	return nil
}

func (p *WgPlugin) syncVppCall() {
	// prepare request
	request := &wg.WgGenkey{}
	// prepare reply
	reply := &wg.WgGenkeyReply{}
	// send request and obtain reply
	err := p.vppChan.SendRequest(request).ReceiveReply(reply)
	if err != nil {
		panic(err)
	}
	// check return value
	if reply.Retval != 0 {
		log.Panicf("Sync call loopback create returned %d", reply.Retval)
	}

	log.Printf("========================================= Private key: %s", reply.PrivateKey)

	request2 := &wg.WgPubkey{
		PrivateKey : reply.PrivateKey,
	}
	// prepare reply
	reply2 := &wg.WgPubkeyReply{}

	err2 := p.vppChan.SendRequest(request2).ReceiveReply(reply2)
	if err2 != nil {
		panic(err2)
	}
	// check return value
	if reply2.Retval != 0 {
		log.Panicf("Sync call loopback create returned %d", reply2.Retval)
	}
	log.Printf("========================================= Public key: %s", reply2.PublicKey)


	request3 := &wg.WgSetDevice{
		PrivateKey : reply.PrivateKey,
		Port: 12312,
	}
	// prepare reply
	reply3 := &wg.WgSetDeviceReply{}

	err3 := p.vppChan.SendRequest(request3).ReceiveReply(reply3)
	if err3 != nil {
		panic(err2)
	}
	// check return value
	if reply3.Retval != 0 {
		log.Panicf("Sync call loopback create returned %d", reply3.Retval)
	}



}

// Close is executed on agent shutdown.
func (p *WgPlugin) Close() error {
	p.vppChan.Close()
	log.Println("Goodbye World!")
	return nil
}

func macParser(mac string) []byte {
	hw, err := net.ParseMAC(mac)
	if err != nil {
		panic(err)
	}
	return hw
}
