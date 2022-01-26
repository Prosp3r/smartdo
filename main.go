package main

import (
	"context"
	"fmt"
	"math/big"

	// "github.com/Prosp3r/smartdo/interact"
	"github.com/Prosp3r/smartdo/interact"
	"github.com/Prosp3r/smartdo/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EvMOSNet = "http://192.168.8.105:8545"
	MainNet  = "https://mainnet.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	GanaChe  = ""

	RinkeByTestNet = "https://rinkeby.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	KovanTestNet   = "https://kovan.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	RopstenTestNet = "https://ropsten.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"

	KeyStoreLocation = "./wallet"
	TestPassword1    = "password"
	TestUserName1    = "prosper"
	TestPassword2    = "password"
	TestUserName2    = "efemena"

	ActiveNet = RopstenTestNet //RinkeByTestNet //KovanTestNet
)

func FailOnError(err error, note string) bool {
	if err != nil {
		fmt.Printf("Error: %v - %v\n", note, err)
		return true
	}
	return false
}

//SampleData - temporary information for testing
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

	//address1//0xd54dBb460e43463D9382E38d06aAf258a27D050a
	//address2//0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d
}

//END SampleData - temporary information for testing

var address1 common.Address //"0xd54dBb460e43463D9382E38d06aAf258a27D050a"
var address2 common.Address //"0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d"

func main() {

	//Load Sample data
	_ = loadSampleData()

	// eClient, err := ethclient.DialContext(context.Background(), ActiveNet)
	eClient, err := ethclient.Dial(ActiveNet)
	_ = FailOnError(err, "Error creating ether client")
	defer eClient.Close()

	/*
		Commands
		1. Check current app balance
		2. Query app
		3. Send tokens to app
		4. Add Task
		5. Change Task
		6. List Tasks
		7. Delete Task
		8. Update Task
		9. Create wallet
	*/

	//Process sample transaction
	// ProcessSampleTransaction(eClient)
	//end process sample transaction

	//Deploy smart Contract to chosen testnet
	// _ = deploy.Deploy(eClient, TestUserName1, TestPassword1)

	//Interact with contract of Hex: 0x4241D10e086895Ca1E08903baB2778e49aa31d37
	TransHex := "0x4241D10e086895Ca1E08903baB2778e49aa31d37"

	// _ = interact.InteractAdd(eClient, TestUserName1, TestPassword1, TransHex)
	_, err = interact.InteractList(eClient, TestUserName1, TestPassword1, TransHex)
	// _ = interact.InteractUpdate(eClient, TestUserName1, TestPassword1, TransHex, "MAKE BURGER", "MAKE MORE BURGER")

	//CheckBalance
	// address, err := utility.GetUserAddress(TestUserName1, TestPassword2)
	// balance := utility.CheckCryptoBalance(*address, eClient)
	// fmt.Printf("Blanace wei: %v\n Balance ETH : %v\n", balance, utility.WeiToEther(balance))

}

func ProcessAddTransaction(eClient *ethclient.Client) {

	senderKeys, err := utility.ReadCryptoKey(TestUserName1, TestPassword1)
	_ = FailOnError(err, "ReadCryptoWallet")

	senderWallet, err := utility.GetUserAddress(TestUserName1, TestPassword1)
	_ = FailOnError(err, "GetUserAddress")

	receiverWallet, err := utility.GetUserAddress(TestUserName2, TestPassword2)
	_ = FailOnError(err, "GetUserAddress")

	//send ether
	amount := big.NewInt(5000000000) //wei
	var AppData []byte = nil
	transaction, err := utility.CreateNewTransaction(*senderWallet, *receiverWallet, amount, eClient, AppData)
	if FailOnError(err, "CreateNewTransaction") == true {
		return
	}

	chainID, err := eClient.NetworkID(context.Background())
	if FailOnError(err, "eClient.NetworkID") == true {
		return
	}
	//sign transaction with private key
	signedTranx, err := types.SignTx(transaction, types.NewEIP155Signer(chainID), senderKeys.PrivateKey)
	if FailOnError(err, "SignTx") == true {
		return
	}

	sendTx, err := utility.SendTransaction(eClient, signedTranx)
	if FailOnError(err, "CreateNewTransaction") == true {
		return
	}

	// sendTxH := common.HexToAddress(*sendTx)
	fmt.Printf("Transaction hash : %v\n\n", *sendTx)

	//TODO
	//Display more information to user eg. Howmuch receiver gets, cost, etc.
	//
}
