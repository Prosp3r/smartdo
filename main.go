package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EvMOSNet = "http://192.168.8.105:8545"
)

func FailOnError(err error, note string) bool {
	if err != nil {
		fmt.Printf("Error: %v - %v\n", note, err)
		return true
	}
	return false
}

func main() {

	eClient, err := ethclient.DialContext(context.Background(), EvMOSNet)
	_ = FailOnError(err, "Error creating ether client")
	defer eClient.Close()

	for {
		block := getLastNetBlock(eClient)
		fmt.Printf("Block Number received => : %v\n", block.Number())
		time.Sleep(time.Second * 5)
	}
}

func getLastNetBlock(ec *ethclient.Client) *types.Block {
	block, err := ec.BlockByNumber(context.Background(), nil)
	_ = FailOnError(err, "-BlockbyNumber")
	return block
}
