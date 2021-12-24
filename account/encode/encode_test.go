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
	address := "5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY"
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
