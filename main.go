package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"

	// "github.com/Prosp3r/smartdo/interact"
	"github.com/Prosp3r/smartdo/deploy"
	// "github.com/Prosp3r/smartdo/interact"
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
	TestUserName1    = "efemena"
	TestPassword2    = "akomeno123,"
	TestUserName2    = "omovie"

	ActiveNet = KovanTestNet //RopstenTestNet //RinkeByTestNet //KovanTestNet
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

	//Run OS flags
	// flagger()

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

	arg := os.Args

	if len(arg) < 1 {
		fmt.Println("Missing options ./smartdo <command> username password \n e.g: ./smartdo adduser username password")
		return
	}
	command := arg[1]
	username := arg[2]
	password := arg[3]

	// fmt.Printf("Command: %v\n Username: %v\n Password: %v\n", command, username, password)

	//Process sample transaction
	// ProcessSampleTransaction(eClient)
	//end process sample transaction

	//Deploy smart contract
	if command == "deploy" {
		//Deploy smart Contract to chosen testnet
		dResult, err := deploy.Deploy(eClient, username, password)
		if FailOnError(err, "Error creating a deployment") == true {
			fmt.Printf("%v\n", err)
			return
		}
		TransHex := dResult.TransactionHex
		fmt.Println(TransHex)
	}

	//Create account
	if command == "adduser" {
		uWallet, err := utility.CreateCryptoWallet(username, password)
		if utility.FailOnError(err, "utility.CreateCryptoWallet") {
			fmt.Println("Cound not create account - ", err)
			return
		}
		fmt.Printf("Your crypto wallet : %v \n Username: %v \n Password: %v \n", uWallet.Address, username, password)
	}

	//Create account
	if command == "mywallet" {
		uCryptoKey, err := utility.ReadCryptoKey(username, password)
		if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
			fmt.Println("Cound not find account - ", err)
			return
		}
		address := utility.GetWalletAddress(uCryptoKey)
		fmt.Printf("Your crypto wallet : %v \n", address)
	}

	if command == "balance" {
		uCryptoKey, err := utility.ReadCryptoKey(username, password)
		if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
			fmt.Println("Cound not find account - ", err)
			return
		}
		address := utility.GetWalletAddress(uCryptoKey)
		balance := utility.CheckCryptoBalance(*address, eClient)
		ethbalance := utility.WeiToEther(balance)

		// fmt.Printf("Your crypto wallet : %v \nBalance: %v\n", address, balance)
		fmt.Printf("Your crypto wallet : %v \nBalance(wei): %v\n Balanace(eth):%f\n", address, balance, ethbalance)
	}

	if command == "sendwei" {
		if len(arg) > 5 {
			recipeient := arg[4]
			sendAmount := arg[5]
			s, err := strconv.Atoi(sendAmount)
			if utility.FailOnError(err, "strconv.Atoi") == true {
				fmt.Println("Cound not read amount to send - ", err)
				return
			}

			amount := new(big.Int).SetUint64(uint64(s))

			// uCryptoKey, err := utility.ReadCryptoKey(username, password)
			// if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
			// 	fmt.Println("Cound not find account - ", err)
			// 	return
			// }
			receiver := common.HexToAddress(recipeient)

			senderAddress, err := utility.GetUserAddress(username, password)
			if utility.FailOnError(err, "utility.GetUserAddress") == true {
				fmt.Println("Could not find sender address")
				return
			}

			senderBalance, err := utility.GetWalletBalance(eClient, senderAddress.Hex())
			fmt.Printf("Sender address: %v \nSender balance: %v \nSending...", senderAddress.Hex(), senderBalance)

			senderKeys, err := utility.ReadCryptoKey(TestUserName1, TestPassword1)
			if FailOnError(err, "ReadCryptoWallet") == true {
				fmt.Printf("Could not fetch sender private key - pleae make sure the password is correct.", err)
				return
			}

			var AppData []byte = nil
			send, err := utility.CreateNewTransaction(*senderAddress, receiver, amount, eClient, AppData)
			if utility.FailOnError(err, "utility.CreateNewTransaction") == true {
				fmt.Println("Cound not send amount - ", err)
				return
			}

			tranxHex, err := utility.BindAndSendTransaction(eClient, send, senderKeys)
			if utility.FailOnError(err, "BindAndSendTransaction") == true {
				fmt.Println("Cound not send amount - ", err)
				return
			}
			fmt.Printf("Trasaction was successful Hex: %v \n\n", *tranxHex)
			return

		}
		fmt.Println("Command sendwei: Missing recipient address")
		return

	}

	if command == "exsendwei" {
		ProcessAddTransaction(eClient, username, password)
	}

	// if command == "cbalance" {
	// 	balance, err := utility.GetWalletBalance(eClient, address1.Hex())
	// 	if utility.FailOnError(err, "") == true {
	// 		fmt.Println("Could no get wallet balance")
	// 		return
	// 	}
	// 	ethbalance := utility.WeiToEther(balance)
	// 	fmt.Printf("Your crypto wallet : %v \nBalance(wei): %v\n Balanace(eth): %v\n", address1, balance, ethbalance)
	// }

	//Interact with contract of Hex: 0x4241D10e086895Ca1E08903baB2778e49aa31d37
	// TransHex := "0x4241D10e086895Ca1E08903baB2778e49aa31d37"

	// _ = interact.InteractAdd(eClient, TestUserName1, TestPassword1, TransHex)
	// _, err = interact.InteractList(eClient, TestUserName1, TestPassword1, TransHex)

	// _ = interact.InteractUpdate(eClient, TestUserName1, TestPassword1, TransHex, "MAKE BURGER", "MAKE MORE BURGER")

	//CheckBalance
	// address, err := utility.GetUserAddress(TestUserName1, TestPassword2)
	// balance := utility.CheckCryptoBalance(*address, eClient)
	// fmt.Printf("Blanace wei: %v\n Balance ETH : %v\n", balance, utility.WeiToEther(balance))

}

func ProcessAddTransaction(eClient *ethclient.Client, fromAccountUsername, fromAccountPassword string) {

	senderKeys, err := utility.ReadCryptoKey(TestUserName1, TestPassword1)
	_ = FailOnError(err, "ReadCryptoWallet")

	senderWallet, err := utility.GetUserAddress(TestUserName1, TestPassword1)
	_ = FailOnError(err, "GetUserAddress")

	receiverWallet, err := utility.GetUserAddress(TestUserName2, TestPassword2)
	_ = FailOnError(err, "GetUserAddress")

	//check
	hashAddToAdd := common.HexToAddress("0x36FcEdA6E0e4044Bb4Dd65D72cC0A9206B83D183")

	//send ether
	amount := big.NewInt(5000000000) //wei
	var AppData []byte = nil
	transaction, err := utility.CreateNewTransaction(*senderWallet, hashAddToAdd, amount, eClient, AppData)
	// transaction, err := utility.CreateNewTransaction(*senderWallet, *receiverWallet, amount, eClient, AppData)
	if FailOnError(err, "CreateNewTransaction") == true {
		return
	}

	//check
	// hashAddToAdd := common.HexToAddress("0x36FcEdA6E0e4044Bb4Dd65D72cC0A9206B83D183")
	fmt.Printf("To Address: %v \n Converted ToAddress: %v \n %v \n\n", transaction.To(), hashAddToAdd, &*receiverWallet)

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
