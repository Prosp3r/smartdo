package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"

	// "github.com/Prosp3r/smartdo/interact"
	"github.com/Prosp3r/smartdo/deploy"
	"github.com/Prosp3r/smartdo/interact"
	"github.com/holiman/uint256"

	// "github.com/Prosp3r/smartdo/interact"
	"github.com/Prosp3r/smartdo/utility"
	"github.com/ethereum/go-ethereum/common"
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
		9. Create wallet
	*/

	arg := os.Args

	if len(arg) < 1 {
		fmt.Println("Missing options ./smartdo <command> username password \n e.g: ./smartdo adduser username password")
		return
	}
	command := arg[1]

	//Deploy smart contract to active network
	if command == "deploy" {

		username := arg[2]
		password := arg[3]

		var Dep utility.DeployedContracts
		//Deploy smart Contract to chosen testnet
		deployname := arg[4]
		dResult, err := deploy.Deploy(eClient, username, password)
		if FailOnError(err, "Error creating a deployment") == true {
			fmt.Printf("%v\n", err)
			return
		}

		Dep.ContractHex = dResult.AddressHex
		Dep.TranxHex = dResult.TransactionHex
		Dep.NetworkDeployed = ActiveNet

		DepJ, err := json.MarshalIndent(Dep, "", " ")
		//
		fmt.Println(Dep.TranxHex)
		fmt.Println(Dep.ContractHex)

		err = ioutil.WriteFile("loadedcontracts/"+deployname, DepJ, 0644)

		//write to file
	}

	//Create account
	if command == "adduser" {
		//Creates a new encrypted ethereum compatible wallet
		username := arg[2]
		password := arg[3]
		uWallet, err := utility.CreateCryptoWallet(username, password)
		if utility.FailOnError(err, "utility.CreateCryptoWallet") {
			fmt.Println("Cound not create account - ", err)
			return
		}
		fmt.Printf("Your crypto wallet : %v \n Username: %v \n Password: %v \n", uWallet.Address, username, password)
	}

	//Print address hex
	if command == "mywallet" {
		// $ ./smartdo mywallet username password
		username := arg[2]
		password := arg[3]
		uCryptoKey, err := utility.ReadCryptoKey(username, password)
		if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
			fmt.Println("Cound not find account - ", err)
			return
		}
		address := utility.GetWalletAddress(uCryptoKey)
		fmt.Printf("Your crypto wallet : %v \n", address)
	}

	//Print address balance
	if command == "balance" {
		// $ ./smartdo mywallet username password
		username := arg[2]
		password := arg[3]

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
		//Send wei to another address
		// $ ./smartdo sendwei username password <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <sendAmount e.g. 200>
		username := arg[2]
		password := arg[3]
		if len(arg) > 5 {
			recipeient := arg[4]
			sendAmount := arg[5]
			s, err := strconv.Atoi(sendAmount)
			if utility.FailOnError(err, "strconv.Atoi") == true {
				fmt.Println("Cound not read amount to send - ", err)
				return
			}

			amount := new(big.Int).SetUint64(uint64(s))

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

	if command == "contract-mint" {
		// $ ./smartdo contract-mint username password <contractName e.g. logi> <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <amountOfTokens e.g. 200>
		//username must be the one that deployed the contract

		username := arg[2] //admin username
		password := arg[3] //admin user password

		contractName := arg[4]     //which contract should be used for this interraction
		recipientAddress := arg[5] //which address will receive the minted tokens

		var ii interact.InputQuery

		//read contract information by name used to save it
		contract, err := utility.ReadStoredContract(contractName)
		if utility.FailOnError(err, "BindAndSendTransaction") == true {
			fmt.Printf("Cound not find smart contract - names: %v\n encountered error %v\n", contractName, err)
			return
		}

		//If required
		sendAmount := arg[6] //amount of tokens to mint to recipient address
		s, err := strconv.Atoi(sendAmount)
		if utility.FailOnError(err, "strconv.Atoi") == true {
			fmt.Println("Cound not read amount to send - ", err)
			return
		}
		ii.Amount = *uint256.NewInt(uint64(s))
		ii.CcontractHex = contract.ContractHex
		ii.Username = username
		ii.Password = password
		ii.AddressRecipient = common.HexToAddress(recipientAddress)

		interact.Mint(eClient, &ii)
		return
	}


	if command == "contract-transfer" {
		// $ ./smartdo contract-transfer username password <contractName e.g. logi> <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <amountOfTokens e.g. 200>
		//username must be the one that deployed the contract

		username := arg[2] //admin username
		password := arg[3] //admin user password

		contractName := arg[4]     //which contract should be used for this interraction
		recipientAddress := arg[5] //which address will receive the minted tokens

		var ii interact.InputQuery

		//read contract information by name used to save it
		contract, err := utility.ReadStoredContract(contractName)
		if utility.FailOnError(err, "BindAndSendTransaction") == true {
			fmt.Printf("Cound not find smart contract - names: %v\n encountered error %v\n", contractName, err)
			return
		}

		//If required
		sendAmount := arg[6] //amount of tokens to mint to recipient address
		s, err := strconv.Atoi(sendAmount)
		if utility.FailOnError(err, "strconv.Atoi") == true {
			fmt.Println("Cound not read amount to send - ", err)
			return
		}
		ii.Amount = *uint256.NewInt(uint64(s))
		ii.CcontractHex = contract.ContractHex
		ii.Username = username
		ii.Password = password
		ii.AddressRecipient = common.HexToAddress(recipientAddress)

		interact.Mint(eClient, &ii)
		return
	}
}
