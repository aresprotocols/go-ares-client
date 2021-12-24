package client

import (
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4"
	"testing"
)

func TestSimpleConnect(t *testing.T) {
	// The following example shows how to instantiate a Substrate API and use it to connect to a node

	url := "wss://rpc.polkadot.io"
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		panic(err)
	}

	chain, err := api.RPC.System.Chain()
	if err != nil {
		panic(err)
	}
	fmt.Println("chain", chain)

	nodeName, err := api.RPC.System.Name()
	if err != nil {
		panic(err)
	}
	fmt.Println("nodeName", nodeName)

	nodeVersion, err := api.RPC.System.Version()
	fmt.Println("nodeVersion", nodeVersion)

	if err != nil {
		panic(err)
	}

	fmt.Printf("You are connected to chain %v using %v v%v\n", chain, nodeName, nodeVersion)
}

func TestSimpleSystem(t *testing.T) {
	// The following example shows how to instantiate a Substrate API and use it to connect to a node

	//url := "wss://rpc.polkadot.io"
	//url := "wss://pub.elara.patract.io/polkadot"
	url := "wss://kusama-rpc.polkadot.io"
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		panic(err)
	}

	health, err := api.RPC.System.Health()
	if err != nil {
		panic(err)
	}
	fmt.Println("Health", health)

	//peers, err := api.RPC.System.Peers()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("peers", peers)

	networkState, err := api.RPC.System.NetworkState()
	fmt.Println("networkState", networkState)

	properties, err := api.RPC.System.Properties()
	fmt.Println("properties", properties)

	if err != nil {
		panic(err)
	}

}

func TestSimpleChain(t *testing.T) {
	// The following example shows how to instantiate a Substrate API and use it to connect to a node

	//url := "wss://rpc.polkadot.io"
	//url := "wss://pub.elara.patract.io/polkadot"
	url := "wss://kusama-rpc.polkadot.io"
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		panic(err)
	}

	hash, err := api.RPC.Chain.GetBlockHashLatest()
	if err != nil {
		panic(err)
	}
	fmt.Println("hash", hash.Hex())

	block, err := api.RPC.Chain.GetBlockLatest()
	if err != nil {
		panic(err)
	}
	fmt.Printf("block %v\n", block.Justification)
	fmt.Printf("block %v\n", block.Block)
	fmt.Printf("Extrinsics %v\n", block.Block.Extrinsics)
	fmt.Printf("Extrinsics %v\n", len(block.Block.Extrinsics))

	fmt.Printf("Header %v\n", block.Block.Header)
	fmt.Printf("Header Digest %v\n", block.Block.Header.Digest)
	fmt.Printf("Header Number %v\n", block.Block.Header.Number)
	fmt.Printf("Header ParentHash %v\n", block.Block.Header.ParentHash.Hex())
	fmt.Printf("Header ExtrinsicsRoot %v\n", block.Block.Header.ExtrinsicsRoot.Hex())
	fmt.Printf("Header StateRoot %v\n", block.Block.Header.StateRoot.Hex())

	//peers, err := api.RPC.System.Peers()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("peers", peers)

	//networkState, err := api.RPC.System.NetworkState()
	//fmt.Println("networkState", networkState)
	//
	//properties, err := api.RPC.System.Properties()
	//fmt.Println("properties", properties)
	//
	//if err != nil {
	//	panic(err)
	//}

}
