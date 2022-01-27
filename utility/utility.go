package utility

import (
	"context"
	"crypto/ecdsa"
	"errors"
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
	// EvMOSNet = "http://192.168.8.105:8545"
	// // MainNet  = "https://mainnet.infura.io/v3/4b0fe94094e047ffa6292fc8065e42b8"
	// MainNet = "https://mainnet.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	// GanaChe = ""

	// RinkByTestNet  = "https://rinkeby.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	// KovanTestNet   = "https://kovan.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
	// RopstenTestNet = "https://ropsten.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"

	KeyStoreLocation = "wallet"
)

func FailOnError(err error, note string) bool {
	if err != nil {
		fmt.Printf("Error: %v - %v\n", note, err)
		return true
	}
	return false
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

func GetUserAddress(username, password string) (*common.Address, error) {

	wallet, err := ReadCryptoKey(username, password)
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
	//ethBalance := WeiToEther(balance)
	return balance
}

//CreateCryptoWallet - Creates an encrypted wallet with the given password
func CreateCryptoWallet(username, password string) (*accounts.Account, error) {

	
	walletLocation := KeyStoreLocation + "/" + username
	_, err := ioutil.ReadDir(walletLocation)
	if err == nil {
		return nil, errors.New("Username already in use")
	}
	
	key := keystore.NewKeyStore(walletLocation, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(password)
	if FailOnError(err, "key.NewAccount") == true {
		return nil, err
	}
	return &account, nil
}

//ReadCryptoKey - Decrypts and returns private key with user's password
func ReadCryptoKey(username, password string) (*keystore.Key, error) {

	walletLocation := KeyStoreLocation + "/" + username

	all, err := ioutil.ReadDir(walletLocation)
	_ = FailOnError(err, "RedDir")
	if len(all) < 1 {
		return nil, err
	}

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
func WeiToEther(weiValue *big.Int) *big.Float {
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
