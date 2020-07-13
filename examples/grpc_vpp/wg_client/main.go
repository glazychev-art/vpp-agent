// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	vpp_wg "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/wg"
	"log"
	"net"
	"sync"
	"time"

	"github.com/namsral/flag"
	"go.ligato.io/cn-infra/v2/agent"
	"go.ligato.io/cn-infra/v2/infra"
	"go.ligato.io/cn-infra/v2/logging/logrus"
	"google.golang.org/grpc"

	"go.ligato.io/vpp-agent/v3/proto/ligato/configurator"
	"go.ligato.io/vpp-agent/v3/proto/ligato/vpp"
	interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
)

var (
	address    = flag.String("address", "172.17.0.3:9111", "address of GRPC server")
	socketType = flag.String("socket-type", "tcp", "socket type [tcp, tcp4, tcp6, unix, unixpacket]")

	dialTimeout = time.Second * 2
)

var exampleFinished = make(chan struct{})

func main() {
	ep := &ExamplePlugin{}
	ep.SetName("wg-client-example")
	ep.Setup()

	a := agent.NewAgent(
		agent.AllPlugins(ep),
		agent.QuitOnClose(exampleFinished),
	)
	if err := a.Run(); err != nil {
		log.Fatal()
	}
}

// ExamplePlugin demonstrates the use of the remoteclient to locally transport example configuration into the default VPP plugins.
type ExamplePlugin struct {
	infra.PluginDeps

	conn *grpc.ClientConn

	wg     sync.WaitGroup
	cancel context.CancelFunc
}

// Init initializes example plugin.
func (p *ExamplePlugin) Init() (err error) {
	// Set up connection to the server.
	p.conn, err = grpc.Dial("unix",
		grpc.WithInsecure(),
		grpc.WithDialer(dialer(*socketType, *address, dialTimeout)),
	)
	if err != nil {
		return err
	}

	client := configurator.NewConfiguratorServiceClient(p.conn)

	// Apply initial VPP configuration.
	go p.demonstrateClient(client)

	// Schedule reconfiguration.
	var ctx context.Context
	ctx, p.cancel = context.WithCancel(context.Background())
	_ = ctx
	/*plugin.wg.Add(1)
	go plugin.reconfigureVPP(ctx)*/

	go func() {
		time.Sleep(time.Second * 35)
		close(exampleFinished)
	}()

	return nil
}

// Close cleans up the resources.
func (p *ExamplePlugin) Close() error {
	logrus.DefaultLogger().Info("Closing example plugin")

	p.cancel()
	p.wg.Wait()

	if err := p.conn.Close(); err != nil {
		return err
	}

	return nil
}

// demonstrateClient propagates snapshot of the whole initial configuration to VPP plugins.
func (p *ExamplePlugin) demonstrateClient(client configurator.ConfiguratorServiceClient) {
	time.Sleep(time.Second * 2)
	p.Log.Infof("Requesting resync..")

	config := &configurator.Config{
		VppConfig: &vpp.ConfigData{
			Interfaces: []*interfaces.Interface{
				memif1, ipip1,
			},
			WgDevice: &vpp_wg.Device {
				PrivateKey: device.PrivateKey,
				Port: device.Port,
			},
			WgPeers: []*vpp_wg.Peer {peer1},
		},
	}
	_, err := client.Update(context.Background(), &configurator.UpdateRequest{
		Update:     config,
		FullResync: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second * 5)
	p.Log.Infof("Requesting get..")

	cfg, err := client.Get(context.Background(), &configurator.GetRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	out, _ := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(cfg)
	fmt.Printf("Config:\n %+v\n", out)

	time.Sleep(time.Second * 5)
	p.Log.Infof("Requesting dump..")

	dump, err := client.Dump(context.Background(), &configurator.DumpRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Dump:\n %+v\n", proto.MarshalTextString(dump))

	p.Log.Infof("Requesting delete..")
	time.Sleep(time.Second * 5)

	delReq := &configurator.DeleteRequest{
		Delete: &configurator.Config {
			VppConfig: &vpp.ConfigData{
				WgPeers: []*vpp_wg.Peer {peer1},
			},
		},
	}

	_, err = client.Delete(context.Background(), delReq)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second * 5)
	p.Log.Infof("Requesting dump..")

	dump, err = client.Dump(context.Background(), &configurator.DumpRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Dump:\n %+v\n", proto.MarshalTextString(dump))
}

// Dialer for unix domain socket
func dialer(socket, address string, timeoutVal time.Duration) func(string, time.Duration) (net.Conn, error) {
	return func(addr string, timeout time.Duration) (net.Conn, error) {
		// Pass values
		addr, timeout = address, timeoutVal
		// Dial with timeout
		return net.DialTimeout(socket, addr, timeoutVal)
	}
}

var (
	memif0 = &vpp.Interface{
		Name:        "memif0",
		Enabled:     true,
		IpAddresses: []string{"10.10.2.1/24"},
		Type:        interfaces.Interface_MEMIF,
		Link: &interfaces.Interface_Memif{
			Memif: &interfaces.MemifLink{
				Id:             1,
				Master:         true,
				Secret:         "secret",
				SocketFilename: "/tmp/memif0.sock",
			},
		},
	}
	memif1 = &vpp.Interface{
		Name:        "memif1",
		Enabled:     true,
		IpAddresses: []string{"10.11.2.1/24"},
		Type:        interfaces.Interface_MEMIF,
		Link: &interfaces.Interface_Memif{
			Memif: &interfaces.MemifLink{
				Id:             1,
				Master:         true,
				Secret:         "secret",
				SocketFilename: "/tmp/memif1.sock",
			},
		},
	}
	ipip0 = &vpp.Interface{
		Name:        "ipip0",
		Enabled:     true,
		IpAddresses: []string{"10.10.3.1/24"},
		Type:        interfaces.Interface_IPIP_TUNNEL,
		Link: &interfaces.Interface_Ipip{
			Ipip: &interfaces.IPIPLink{
				TunnelMode: 0,
				SrcAddr: "10.10.2.1",
				DstAddr: "10.10.2.2",
			},
		},
	}
	ipip1 = &vpp.Interface{
		Name:        "ipip1",
		Enabled:     true,
		IpAddresses: []string{"10.11.3.1/24"},
		Type:        interfaces.Interface_IPIP_TUNNEL,
		Link: &interfaces.Interface_Ipip{
			Ipip: &interfaces.IPIPLink{
				TunnelMode: 0,
				SrcAddr: "10.11.2.1",
				DstAddr: "10.11.2.2",
			},
		},
	}

	device = &vpp.WgDevice{
		PrivateKey: "gIjXzrQfIFf80d0O8Hd2KhcfkKLRncc+8C70OjotIW8=",
		Port: 12312,
	}
	peer0 = &vpp.WgPeer{
		PublicKey: "33GyVvUQalLCscTfN8TxtTp/ixtSWg55PhHy0aWABHQ=",
		Port: 12314,
		Endpoint: "10.10.2.2",
		AllowedIp: "10.10.3.2",
		TunInterface: "ipip0",
	}
	peer1 = &vpp.WgPeer{
		PublicKey: "dE30h1yk5hS4lPjl5Vpwa8YAdouJhdDgHnrw9WjzpFI=",
		Port: 12314,
		Endpoint: "10.11.2.2",
		AllowedIp: "10.11.3.2",
		TunInterface: "ipip1",
		PersistentKeepalive: 15,
	}
)