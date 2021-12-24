package account

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vedhavyas/go-subkey"
	"github.com/vedhavyas/go-subkey/sr25519"
	"polkadot/ares/account/encode"
	"testing"
)

const (
	// DevPhrase is default phrase used for dev test accounts
	DevPhrase = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

	junctionIDLen = 32
)

func TestSr25519(t *testing.T) {
	c := struct {
		uri       string
		seed      string
		publicKey string
		accountID string
		ss58Addr  string
		network   uint8
		err       bool
	}{}

	c.uri = "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap"
	c.seed = "0x18446f2d685492c3086391aabe8f5e235c3c2e02521985650f0c97052237e717"
	c.publicKey = "0x88af895626c47cf1235ec3898d238baeb41adca3117b9a77bc2f6b78eca0771b"
	c.ss58Addr = "146Df8yer1y9PyxiDmqV3NXqwHq2ahLb3Jkg1zPETRrkaSwb"
	c.network = 0

	s, err := subkey.DeriveKeyPair(sr25519.Scheme{}, c.uri)
	if err != nil {
		assert.True(t, c.err)
		return
	}

	pub := s.Public()
	assert.Equal(t, c.publicKey, subkey.EncodeHex(pub))
	if c.accountID != "" {
		assert.Equal(t, c.accountID, subkey.EncodeHex(s.AccountID()))
	}
	fmt.Println("seed", len(s.Seed()))
	seed := subkey.EncodeHex(s.Seed())
	if s.Seed() == nil {
		seed = ""
	}
	assert.Equal(t, c.seed, seed)
	gotSS58Addr, err := s.SS58Address(c.network)

	address := encode.EncodeAddress(s.AccountID(), 0)
	fmt.Println("address", address, "gotSS58Addr", gotSS58Addr)

	fmt.Println("gotSS58Addr", gotSS58Addr)
	assert.NoError(t, err)
	assert.Equal(t, c.ss58Addr, gotSS58Addr)
	msg := []byte("msg")
	sig, err := s.Sign(msg)
	fmt.Println("sig", len(sig))
	assert.NoError(t, err)
	assert.True(t, s.Verify(msg, sig))
}

func TestKeyInfo(t *testing.T) {
	c := struct {
		uri       string
		seed      string
		publicKey string
		accountID string
		ss58Addr  string
		network   uint8
		err       bool
	}{}

	c.uri = "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap"
	c.seed = "0x18446f2d685492c3086391aabe8f5e235c3c2e02521985650f0c97052237e717"
	c.publicKey = "0x88af895626c47cf1235ec3898d238baeb41adca3117b9a77bc2f6b78eca0771b"
	c.ss58Addr = "146Df8yer1y9PyxiDmqV3NXqwHq2ahLb3Jkg1zPETRrkaSwb"
	c.network = 0

	s, err := subkey.DeriveKeyPair(sr25519.Scheme{}, c.uri)
	if err != nil {
		assert.True(t, c.err)
		return
	}

	pub := s.Public()
	fmt.Println("pub", subkey.EncodeHex(pub))
	fmt.Println("accountID", subkey.EncodeHex(s.AccountID()))
	fmt.Println("seed", len(s.Seed()))
	seed := subkey.EncodeHex(s.Seed())
	fmt.Println("seed", seed)

	gotSS58Addr, err := s.SS58Address(c.network)
	fmt.Println("gotSS58Addr", gotSS58Addr)
}

func TestDevAccount(t *testing.T) {
	uris := []string{
		DevPhrase + "/Alice", DevPhrase + "//Alice",
	}
	network := uint8(42)
	for _, uri := range uris {
		s, err := subkey.DeriveKeyPair(sr25519.Scheme{}, uri)
		if err != nil {
			return
		}

		pub := s.Public()
		fmt.Println("pub", subkey.EncodeHex(pub))
		fmt.Println("accountID", subkey.EncodeHex(s.AccountID()))
		fmt.Println("seed", len(s.Seed()))
		seed := subkey.EncodeHex(s.Seed())
		fmt.Println("seed", seed)

		gotSS58Addr, err := s.SS58Address(network)
		fmt.Println("gotSS58Addr", gotSS58Addr)
	}
}

//var TestKeyringPairAlice = KeyringPair{
//	URI:       "//Alice",
//	PublicKey: []byte{0xd4, 0x35, 0x93, 0xc7, 0x15, 0xfd, 0xd3, 0x1c, 0x61, 0x14, 0x1a, 0xbd, 0x4, 0xa9, 0x9f, 0xd6, 0x82, 0x2c, 0x85, 0x58, 0x85, 0x4c, 0xcd, 0xe3, 0x9a, 0x56, 0x84, 0xe7, 0xa5, 0x6d, 0xa2, 0x7d}, //nolint:lll
//	Address:   "5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY",
//}
