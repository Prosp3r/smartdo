package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "golang.org/x/tools/cmd/stringer"
)

var (
	EvMOSNet = "http://192.168.8.105:8545"
	// MainNet  = "https://mainnet.infura.io/v3/4b0fe94094e047ffa6292fc8065e42b8"
	MainNet = "https://mainnet.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	GanaChe = ""

	RinkByTestNet  = "https://rinkeby.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	KovanTestNet   = "https://kovan.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	RopstenTestNet = "https://ropsten.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"

	KeyStoreLocation = "./wallet"
	TestPassword1    = "password"
	TestUserName1    = "prosper"
	TestPassword2    = "password"
	TestUserName2    = "efemena"

	ActiveNet = RopstenTestNet //RinkByTestNet //KovanTestNet
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

	//Process sample transaction
	ProcessSampleTransaction(eClient)
	//end process sample transaction

}

func ProcessSampleTransaction(eClient *ethclient.Client) {

	senderKeys, err := ReadCryptoKey(TestPassword1, TestUserName1)
	_ = FailOnError(err, "ReadCryptoWallet")

	senderWallet, err := GetUserAddress(TestPassword1, TestUserName1)
	_ = FailOnError(err, "GetUserAddress")

	receiverWallet, err := GetUserAddress(TestPassword2, TestUserName2)
	_ = FailOnError(err, "GetUserAddress")

	//send ether
	amount := big.NewInt(5000000000) //wei
	var AppData []byte = nil
	transaction, err := CreateNewTransaction(*senderWallet, *receiverWallet, amount, eClient, AppData)
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

	sendTx, err := SendTransaction(eClient, signedTranx)
	if FailOnError(err, "CreateNewTransaction") == true {
		return
	}

	// sendTxH := common.HexToAddress(*sendTx)
	fmt.Printf("Transaction hash : %v\n\n", *sendTx)

	//TODO
	//Display more information to user eg. Howmuch receiver gets, cost, etc.
	//
}

//SendTransaction - sends transaction to blockchain
func SendTransaction(eClient *ethclient.Client, tranx *types.Transaction) (*string, error) {
	err := eClient.SendTransaction(context.Background(), tranx)
	if FailOnError(err, "SendTransaction") == true {
		return nil, err
	}

	TxHash := tranx.Hash().Hex()
	return &TxHash, nil
}

func GetUserAddress(password, username string) (*common.Address, error) {

	wallet, err := ReadCryptoKey(password, username)
	if FailOnError(err, "ReadCryptoWallet") == true {
		return nil, err
	}

	cryptoAddress := GetWalletAddress(wallet)
	return cryptoAddress, nil
}

//CreateNewTransaction - Creates a new transaction compatible on the network given in client connection
func CreateNewTransaction(fromAddress, toAddress common.Address, amount *big.Int, client *ethclient.Client, AppData []byte) (*types.Transaction, error) {
	var gassLimit uint64 = 21000

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if FailOnError(err, "client.SuggestGasPrice") == true {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if FailOnError(err, "client.PendingNonceAt") == true {
		return nil, err
	}

	// //:::::::::check if sender balance is enough to make payment
	// senderBalance := CheckCryptoBalance(fromAddress, client)

	// total := gasPrice.Add(amount, gasPrice)
	// // fmt.Printf("Total to send wei: %v\n", total)
	// fmt.Printf("Current sender balance: %v\n Needed amount : %v\n", weiToEther(senderBalance), weiToEther(total))

	// if total.CmpAbs(senderBalance) <= 0 {
	// 	return nil, errors.New("Insufficient funds to initiate transaction")
	// }
	// //:::END::::::check if sender balance is enough to make payment
	trx := types.NewTransaction(nonce, toAddress, amount, gassLimit, gasPrice, AppData)
	// fmt.Printf("Transaction: %v\n To address %v\n From address %v\n", trx, toAddress, fromAddress)
	return trx, nil
}

//CheckCryptoBalance - returns the balance of the given address on the ethereum(mainor testnet) network
func CheckCryptoBalance(walletAddress common.Address, eClient *ethclient.Client) *big.Int {

	balance, err := eClient.BalanceAt(context.Background(), walletAddress, nil)
	_ = FailOnError(err, "eClient.BalanceAt")
	//ethBalance := weiToEther(balance)
	return balance
}

//CreateCryptoWallet - Creates an encrypted wallet with the given password
func CreateCryptoWallet(password, username string) *accounts.Account {
	walletLocation := KeyStoreLocation + "/" + username
	key := keystore.NewKeyStore(walletLocation, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(password)
	_ = FailOnError(err, "key.NewAccount")
	fmt.Println(account.Address)
	return &account
}

//ReadCryptoKey - Decrypts and returns private key with user's password
func ReadCryptoKey(password, username string) (*keystore.Key, error) {

	walletLocation := KeyStoreLocation + "/" + username
	all, err := ioutil.ReadDir(walletLocation)
	_ = FailOnError(err, "RedDir")
	walletFile := all[0]

	readFile, err := ioutil.ReadFile(walletLocation + "/" + walletFile.Name())
	if FailOnError(err, "ioutil.ReadFile") == true {
		return nil, err
	}
	key, err := keystore.DecryptKey(readFile, password)
	_ = FailOnError(err, "keystore.DecryptKey")

	// cryptoAddress := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	return key, nil
}

//GetWalletAddress - Returns wallet address
func GetWalletAddress(privKey *keystore.Key) *common.Address {
	walletHex := crypto.PubkeyToAddress(privKey.PrivateKey.PublicKey).Hex()
	cryptoAddress := common.HexToAddress(walletHex)
	return &cryptoAddress
}

//getLastNetBlock - Returns the last block minned on the network
func getLastNetBlockWei(ec *ethclient.Client) *types.Block {
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

//weiToEther - converts given wei to ether
func weiToEther(weiValue *big.Int) *big.Float {
	//11 eth = 10^18 wei
	var fB = new(big.Float)
	fB.SetString(weiValue.String())
	ethValue := new(big.Float).Quo(fB, big.NewFloat(math.Pow10(18)))

	return ethValue
}

// //weiToEther - converts given wei to ether
// func EtherTowei(EtherValue *big.Float) *big.Int {
// 	return new(big.Int).Mul(EtherValue, big.NewInt(Ether))
// }

//GenPrivateKey - will generate and return a pointer to a new ECDSA private key
func GenPrivateKey() (*ecdsa.PrivateKey, error) {
	pvk, err := crypto.GenerateKey()
	fe := FailOnError(err, "GenPrivateKey")
	if fe == true {
		return nil, err
	}
	return pvk, nil
}

//PrivateKeyTostring - converts given ecdsa private key to human readable string
func PrivateKeyTostring(pvk *ecdsa.PrivateKey) (*string, error) {
	pData := crypto.FromECDSA(pvk)
	kee := string(hexutil.Encode(pData))
	return &kee, nil
}

//GenPublicKey - will generate and a return pointer to a new ECDSA public key
func GenPublicKey(pvk *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &pvk.PublicKey
}

//PrivateKeyTostring - converts given ecdsa private key to human readable string
func PublicKeyTostring(pvk *ecdsa.PublicKey) *string {
	pData := crypto.FromECDSAPub(pvk)
	kee := string(hexutil.Encode(pData))
	return &kee
}

//PubKeyToAddress - return a pointer to a crypto address hex string
func PubKeyToAddress(kee *ecdsa.PublicKey) *string {
	addy := crypto.PubkeyToAddress(*kee).Hex()
	return &addy
}
