package main

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func CreatePrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, &chaincfg.MainNetParams, true)
}

func GetAddressPubKey(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), &chaincfg.MainNetParams)
}

func GenAddressByPubKey(pubKey *btcutil.AddressPubKey) (string, error) {
	witnessProg := btcutil.Hash160(pubKey.PubKey().SerializeCompressed())
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	address := addressWitnessPubKeyHash.EncodeAddress()

	return address, nil
}
