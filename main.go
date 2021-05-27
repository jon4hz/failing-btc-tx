package main

import (
	"fmt"
	"log"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/jon4hz/malicious-btc-tx/config"
)

func main() {

	config.Load()
	cfg := config.Get()

	ntfnHandlers := rpcclient.NotificationHandlers{
		OnAccountBalance: func(account string, balance btcutil.Amount, confirmed bool) {
			log.Printf("New balance for account %s: %v", account,
				balance)
		},
		OnRescanProgress: func(hash *chainhash.Hash, height int32, blkTime time.Time) {
			log.Println("rescan in progress")
		},
		OnRescanFinished: func(hash *chainhash.Hash, height int32, blkTime time.Time) {
			log.Println("rescan finished")
		},
	}

	connCfg := &rpcclient.ConnConfig{
		Host:       fmt.Sprintf("%s:28332", cfg.RPCHost),
		Endpoint:   "ws",
		User:       cfg.RPCUser,
		Pass:       cfg.RPCPassword,
		DisableTLS: true,
	}

	client, err := rpcclient.New(connCfg, &ntfnHandlers)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(client.Connect(1))
}
