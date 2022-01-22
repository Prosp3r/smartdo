package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EvMOSNet      = "http://192.168.8.105:8545"
	MainNet       = "https://mainnet.infura.io/v3/4b0fe94094e047ffa6292fc8065e42b8"
	GanaChe       = ""
	RinkByTestNet = "https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
)

func FailOnError(err error, note string) bool {
	if err != nil {
		fmt.Printf("Error: %v - %v\n", note, err)
		return true
	}
	return false
}

type SampleData struct {
	address         string
	contractAddress string
	tranxHash       string
}

var SD SampleData

func loadSampleData() bool {
	SD.address = "0x9e77cc237460bbbc8935457e487d4ecfa59030c3"
	SD.contractAddress = "0x4b978d499f2ae9be60e765b7c531faf847863255"
	SD.tranxHash = "0xf943f16dc36ad99ecce29c3eb351a7ff744bf86093dace022ed5fbffbb651af9"

	return true
}

func main() {

	//Load Sample data
	_ = loadSampleData()

	eClient, err := ethclient.DialContext(context.Background(), RinkByTestNet)
	_ = FailOnError(err, "Error creating ether client")
	defer eClient.Close()

	for {
		block := getLastNetBlock(eClient)

		fmt.Printf("Block Number received => : %v \nTransactions => %v\n", block.Number(), block.Transactions().Len())
		js, err := block.Header().MarshalJSON()
		_ = FailOnError(err, "JSON Conversion")

		jb := block.Body()
		// _ = FailOnError(err, "JSON Conversion")

		fmt.Printf("Block Header : %v\n\n Block Body :%v\n\n", string(js), jb)

		// addr := "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
		bal, err := getWalletBalance(eClient, SD.address)
		_ = FailOnError(err, "getWalletBalance")
		fmt.Printf("WalletBalance : => %v \n", bal)
		time.Sleep(time.Second * 5)
	}
}

//getLastNetBlock - Returns the last block minned on the network
func getLastNetBlock(ec *ethclient.Client) *types.Block {
	block, err := ec.BlockByNumber(context.Background(), nil)
	_ = FailOnError(err, "-getLastNetBlock")
	return block
}

//getWalletBalance - Returns the decimal balance of give wallet hex string
func getWalletBalance(ec *ethclient.Client, addr string) (*big.Int, error) {
	address := common.HexToAddress(addr)
	fmt.Printf("Address: %v - ", address)
	balance, err := ec.BalanceAt(context.Background(), address, nil)
	if FailOnError(err, "getWalletBalance") == true {
		return nil, err
	}
	return balance, nil
}
