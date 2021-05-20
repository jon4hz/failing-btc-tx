package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
)

func main() {
	fmt.Println("Let's go...")

	ntfnHandlers := rpcclient.NotificationHandlers{
		OnAccountBalance: func(account string, balance btcutil.Amount, confirmed bool) {
			log.Printf("New balance for account %s: %v", account,
				balance)
		},
	}

	connCfg := &rpcclient.ConnConfig{
		Host:     "localhost:18332",
		Endpoint: "ws",
		User:     "yourrpcuser",
		Pass:     "yourrpcpass",
	}

	client, err := rpcclient.New(connCfg, &ntfnHandlers)
	_, _ = client, err

}
