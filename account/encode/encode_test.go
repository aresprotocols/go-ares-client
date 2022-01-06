package encode

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"log"
	"polkadot/ares/account/bip39"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	address := "4Smy4MPq2ER7ioEwCgGt2NEKaz79MbQfmNQRp8guxcgkYnaC"
	fmt.Println("len", len([]byte(address)))
	publicKey, ss58Prefix, err := DecodeAddress(address)
	if err != nil {
		log.Fatalf("Failed Get create: %s", err.Error())
	}
	data := [][]string{
		{"Address", fmt.Sprintf("%s", address)},
		{"PublicKey", fmt.Sprintf("0x%x", publicKey)},
		{"ss58Prefix", fmt.Sprintf("%d", ss58Prefix)},
	}
	fmt.Println("data ", data)
}

func TestED25519(t *testing.T) {
	var (
		privateKey []byte
		publicKey  []byte
	)

	//mnemonic, entropy, err := bip39.GenerateMnemonic()
	//if err != nil {
	//	log.Fatalf("Failed GenerateMnemonic: %s", err.Error())
	//}

	mnemonic := strings.Split("gown away walk sword actress fish memory neutral avoid network alien choice", " ")
	entropy, err := hex.DecodeString("65220fda6e202caf62bca71012981914")
	if err != nil {
		log.Fatalf("Failed GenerateMnemonic: %s", err.Error())
	}
	fmt.Println("mnemonic", mnemonic, "mnemonic", mnemonic[0])

	seed, err := bip39.MnemonicToSeed(strings.Join(mnemonic, " "), "m7436528")
	if err != nil {
		log.Fatalf("Failed MnemonicToSeed: %s", err.Error())
	}
	extendedPrivateKey := ed25519.NewKeyFromSeed(seed[:32])
	privateKey = extendedPrivateKey[:32]
	publicKey = extendedPrivateKey[32:]
	data := [][]string{
		{"entropy", fmt.Sprintf("%x", entropy)},
		{"mnemonic", fmt.Sprintf("%s", strings.Join(mnemonic, " "))},
	}

	fmt.Println("data", data)

	if err != nil {
		log.Fatalf("Failed Get ss58Prefix: %s", err.Error())
	}
	address := EncodeAddress(publicKey, 37)

	data = append(data, [][]string{
		{"PrivateKey", fmt.Sprintf("0x%x", privateKey)},
		{"PublicKey", fmt.Sprintf("0x%x", publicKey)},
		{"Address", fmt.Sprintf("%s", address)},
	}...)
	fmt.Println("data", data)

}

func TestPub(t *testing.T) {
	dataStr := "58ccb645829bad32a700595d74246c16bf0b981b23367d638c5cef7d31860b65e9bf2c72bc5d46e0d11c57e68b16ac087b487497495984e19a05a6a317da064edc818436afabf30dd0b523ee8440160bc97b25e264d1d054d7d6d342b9ecf353"
	data, err := hex.DecodeString(dataStr)
	aura := data[:32]
	grand := data[32:64]
	ares := data[64:]

	fmt.Println("data", len(data), " err ", err)
	fmt.Println("aura", EncodeAddress(aura, 34), " pub ", hex.EncodeToString(aura))
	fmt.Println("grand", EncodeAddress(grand, 34), " pub ", hex.EncodeToString(grand))
	fmt.Println("ares", EncodeAddress(ares, 34), " pub ", hex.EncodeToString(ares))

	address := "4Rma6tk2UtaGxPh5dozPz4DVFM2R4wK4iZtXP4kXzKqMv27F"
	publicKey, ss58Prefix, err := DecodeAddress(address)
	if err != nil {
		log.Fatalf("Failed Get create: %s", err.Error())
	}
	dataMap := [][]string{
		{"Address", fmt.Sprintf("%s", address)},
		{"PublicKey", fmt.Sprintf("0x%x", publicKey)},
		{"ss58Prefix", fmt.Sprintf("%d", ss58Prefix)},
	}
	fmt.Println("data ", dataMap)
	//     stash
	// 1   4SJT3cozQ7Uv31M8A1q5ysarUEtv58xcoA5GgWBnoZ3b7G5w
	// 2   4UHtW2qVT6A993ViBE7hwe4oXG4UX19bmEedw41rvatJjeWC
	// 3   4TfqZ8mc4FpbNF66Qh73dyisEPYYxoY5mn7NE18LqX2AqgLT
	// 4   4RTJuWG29fQKBU8rr3kAc27rTHyzNts6gsqVkJKrrkp18cfb
	// 5   4Rma6tk2UtaGxPh5dozPz4DVFM2R4wK4iZtXP4kXzKqMv27F

	//auth	0x71cb1378dc818436afabf30dd0b523ee8440160bc97b25e264d1d054d7d6d342
	//block_number	1408494777
	//pre_check_auth	0x851df65708ecdc14e2dd427724c60c6879a1aeade21d9708c30c4477f679dde9
	//pre_check_stash	0x4c30010070214e02fb2ec155a4c7bb8c122864b3b03f58c4ac59e8d83af7dc29

	//pre_check_stash: 4RVKgtvDAitcfeYwkrbK2jXRE1d9VJypheMSp9jv3UfZLhK6
	//pre_check_auth: 4Smy4MPq2ER7ioEwCgGt2NEKaz79MbQfmNQRp8guxcgkYnaC
	//auth: 4SLdXVE5Rn4YMsCM6XaaqctKKJbkBioLtkKZMet1bFf4RW2v
	//block_number: 1,408,494,777

	dataStr = "71cb1378dc818436afabf30dd0b523ee8440160bc97b25e264d1d054d7d6d342"
	data, err = hex.DecodeString(dataStr)
	fmt.Println("aura", EncodeAddress(data, 34))
	dataStr = "851df65708ecdc14e2dd427724c60c6879a1aeade21d9708c30c4477f679dde9"
	data, err = hex.DecodeString(dataStr)
	fmt.Println("aura", EncodeAddress(data, 34))
	dataStr = "4c30010070214e02fb2ec155a4c7bb8c122864b3b03f58c4ac59e8d83af7dc29"
	data, err = hex.DecodeString(dataStr)
	fmt.Println("aura", EncodeAddress(data, 34))
}
