package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

var (
	utxo         = "5493f59f49fff5ed945d17a716fa7e6d72e4c5285c0a8d40adc9a45381cf0940"
	pubkeyScript = "0014a6d1ca1be6eae2547164349e5b8b4df019708f31"
	balance      = 50000
	dest         = "bc1q5mgu5xlxat39gutyxj09hz6d7qvhpre3qemw6d"
)

// https://live.blockcypher.com/btc-testnet/pushtx/

func main() {
	prvKey, err := CreatePrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := CreateTx(prvKey.String(), dest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
}

func NewTx() (*wire.MsgTx, error) {
	return wire.NewMsgTx(wire.TxVersion), nil
}

func CreateTx(privKey string, destination string) (string, error) {

	destinationAddr, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	destinationAddrByte, err := txscript.PayToAddrScript(destinationAddr)
	if err != nil {
		return "", err
	}

	redeemTx, err := NewTx()
	if err != nil {
		return "", err
	}

	utxoHash, err := chainhash.NewHashFromStr(utxo)
	if err != nil {
		return "", err
	}

	outPoint := wire.NewOutPoint(utxoHash, 0)

	txIn := wire.NewTxIn(outPoint, nil, nil)
	redeemTx.AddTxIn(txIn)

	redeemTxOut := wire.NewTxOut(int64(balance), destinationAddrByte)
	redeemTx.AddTxOut(redeemTxOut)

	finalRawTx, err := SignTx(privKey, pubkeyScript, redeemTx)
	if err != nil {
		return "", err
	}

	return finalRawTx, nil
}

func SignTx(privKey string, pkScript string, redeemTx *wire.MsgTx) (string, error) {

	wif, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return "", err
	}

	sourcePKScript, err := hex.DecodeString(pkScript)
	if err != nil {
		return "", nil
	}

	signature, err := txscript.SignatureScript(redeemTx, 0, sourcePKScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return "", nil
	}

	redeemTx.TxIn[0].SignatureScript = signature

	var signedTx bytes.Buffer
	redeemTx.Serialize(&signedTx)

	hexSignedTx := hex.EncodeToString(signedTx.Bytes())

	return hexSignedTx, nil
}
