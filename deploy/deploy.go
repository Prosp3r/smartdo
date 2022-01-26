package deploy

import (
	"context"
	"fmt"
	"math/big"

	todo "github.com/Prosp3r/smartdo/gen"
	"github.com/Prosp3r/smartdo/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DeployResult struct {
	AddressHex     string `json:"addresshex"`
	TransactionHex string `json:"transactionhex"`
}

var Result DeployResult

func Deploy(con *ethclient.Client, username, password string) (*DeployResult, error) {
	//Get user's crypto wallet
	userCryptoAddy, err := utility.ReadCryptoKey(password, username)
	if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
		return nil, err
	}
	//Get the nonce dynamically from the network
	nonce, err := con.PendingNonceAt(context.Background(), userCryptoAddy.Address)
	if utility.FailOnError(err, "con.PendingNonceAt") == true {
		return nil, err
	}
	//get estimated gas price
	gasPrince, err := con.SuggestGasPrice(context.Background())
	if utility.FailOnError(err, "con.SuggestGasPrice") == true {
		return nil, err
	}
	//Get the network chain id
	chainID, err := con.NetworkID(context.Background())
	if utility.FailOnError(err, "con.NetworkID") == true {
		return nil, err
	}
	//Binding
	auth, err := bind.NewKeyedTransactorWithChainID(userCryptoAddy.PrivateKey, chainID)
	if utility.FailOnError(err, "bind.NewKeyedTransactorWithChainID") == true {
		return nil, err
	}

	auth.GasPrice = gasPrince
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))
	//trigger a deployment
	a, tx, _, err := todo.DeployTodo(auth, con)
	if utility.FailOnError(err, "todo.DeployTodo") == true {
		return nil, err
	}

	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

	Result.AddressHex = a.Hex()
	Result.TransactionHex = tx.Hash().Hex()

	return &Result, nil
}
