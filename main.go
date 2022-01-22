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
)

var (
	EvMOSNet         = "http://192.168.8.105:8545"
	MainNet          = "https://mainnet.infura.io/v3/4b0fe94094e047ffa6292fc8065e42b8"
	GanaChe          = ""
	RinkByTestNet    = "https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	KeyStoreLocation = "./wallet"
	TestPassword     = "password"
	TestUserName     = "user_name"
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
}

//END SampleData - temporary information for testing

func main() {

	//Load Sample data
	_ = loadSampleData()

	eClient, err := ethclient.DialContext(context.Background(), RinkByTestNet)
	_ = FailOnError(err, "Error creating ether client")
	defer eClient.Close()

	// a := CreateCryptoWallet(TestPassword, TestUserName)
	// fmt.Println(a.Address)
	// fmt.Println(a.URL)

	walletPrvKey, err := ReadCryptoWallet(TestPassword, TestUserName)
	_ = FailOnError(err, "ReadCryptoWallet")
	cryptoAddress := crypto.PubkeyToAddress(walletPrvKey.PrivateKey.PublicKey).Hex()
	fmt.Println(cryptoAddress)

	// for {
	// 	//1. block := getLastNetBlock(eClient)
	// 	//2. bal, err := getWalletBalance(eClient, SD.address)
	// 	// _ = FailOnError(err, "getWalletBalance")
	// 	// fmt.Printf("\n\nWalletBalance : => %v \n ETH Value : %v\n", bal, weiToEther(bal))

	// 	pk, err := GenPrivateKey()
	// 	_ = FailOnError(err, "GenPrivateKey")
	// 	prKString, err := PrivateKeyTostring(pk)
	// 	_ = FailOnError(err, "PrivateKeyToString")

	// 	pubKey := GenPublicKey(pk)
	// 	puKString := PublicKeyTostring(pubKey)
	// 	//generate address hex
	// 	cryptoAddress := PubKeyToAddress(pubKey)

	// 	fmt.Printf("\n\nPrivateKey: %v\n PublicKey: %v\n Addres: %v\n", *prKString, *puKString, *cryptoAddress)

	// 	time.Sleep(time.Second * 5)
	// }
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

func ReadCryptoWallet(password, username string) (*keystore.Key, error) {

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

	// privateData := crypto.FromECDSA(key.PrivateKey)
	// privateKeyString := hexutil.Encode(privateData)

	// publicData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	// publicKeyString := hexutil.Encode(publicData)
	cryptoAddress := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println(cryptoAddress)
	return key, nil
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

func weiToEther(weiValue *big.Int) *big.Float {
	//11 eth = 10^18 wei
	var fB = new(big.Float)
	fB.SetString(weiValue.String())
	ethValue := new(big.Float).Quo(fB, big.NewFloat(math.Pow10(18)))

	return ethValue
}

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
