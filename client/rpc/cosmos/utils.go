package cosmos

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/rs/zerolog/log"
)

func CreateAccountFromMnemonic(mnemonic string) (*secp256k1.PrivKey, types.AccAddress, error) {
	// Derive the seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Create master key and derive the private key
	// Using "m/44'/118'/0'/0/0" for Cosmos
	master, ch := hd.ComputeMastersFromSeed(seed)
	privKeyBytes, err := hd.DerivePrivateKeyForPath(master, ch, "m/44'/118'/0'/0/0")
	if err != nil {
		return nil, nil, err
	}

	// Create private key and get address
	privKey := &secp256k1.PrivKey{Key: privKeyBytes}
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("scalar", "scalarvaloper")
	addr := types.AccAddress(privKey.PubKey().Address())
	log.Debug().Msgf("Created account with address: %s from mnemonic: %s", addr.String(), mnemonic)
	return privKey, addr, nil
}
