package client

import (
	"fmt"
	"github.com/akamensky/base58"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"golang.org/x/crypto/blake2b"
	"math/big"
	"testing"
)

func TestSession(t *testing.T) {
	// This sample shows how to create a transaction to make a transfer from one an account to another.

	// Instantiate the API
	api, err := gsrpc.NewSubstrateAPI("wss://gladios.aresprotocol.io")
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	key1, err := types.CreateStorageKey(meta, "Session", "Validators", nil)
	var validators []types.AccountID
	ok, err := api.RPC.State.GetStorageLatest(key1, &validators)
	if err != nil || !ok {
		fmt.Println("err ", err)
		panic(err)
	}
	fmt.Printf("Current validators:\n")
	for i, v := range validators {
		var raw []byte
		raw = append([]byte{byte(34)}, v[:]...)
		prefix := []byte("SS58PRE")
		checksum := blake2b.Sum512(append(prefix, raw...))
		address := base58.Encode(append(raw, checksum[0:2]...))

		key, err := types.CreateStorageKey(meta, "System", "Account", v[:])
		if err != nil {
			panic(err)
		}

		var accountInfo types.AccountInfo
		ok, err = api.RPC.State.GetStorageLatest(key, &accountInfo)
		if err != nil || !ok {
			fmt.Println("err ", err)
			panic(err)
		}
		previous := accountInfo.Data.Free
		fmt.Printf("\tValidator %v: %#x  %s  %v %d \n", i, v, address, previous.Int, toAres(previous.Int).Int64())

		//_, err := CreateStorageKey(m, "Staking", "ErasStakers")
		//key, err := CreateStorageKey(m, "Session", "NextKeys",
	}
}

func TestSessionIndex(t *testing.T) {
	api, err := gsrpc.NewSubstrateAPI("wss://gladios.aresprotocol.io")
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	key, err := types.CreateStorageKey(meta, "Session", "CurrentIndex", nil)
	var index types.U32
	ok, err := api.RPC.State.GetStorageLatest(key, &index)
	if err != nil || !ok {
		fmt.Println("err ", err)
		panic(err)
	}
	fmt.Println("sessionIndex ", index)
}

func toAres(val *big.Int) *big.Int {
	baseUnit := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Quo(new(big.Int).Set(val), baseUnit)
}
