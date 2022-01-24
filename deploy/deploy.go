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

func Deploy(con *ethclient.Client, username, password string) bool {
	//Get user's crypto wallet
	userCryptoAddy, err := utility.ReadCryptoKey(password, username)
	if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
		return false
	}

	nonce, err := con.PendingNonceAt(context.Background(), userCryptoAddy.Address)
	if utility.FailOnError(err, "con.PendingNonceAt") == true {
		return false
	}

	gasPrince, err := con.SuggestGasPrice(context.Background())
	if utility.FailOnError(err, "con.SuggestGasPrice") == true {
		return false
	}

	chainID, err := con.NetworkID(context.Background())
	if utility.FailOnError(err, "con.NetworkID") == true {
		return false
	}

	auth, err := bind.NewKeyedTransactorWithChainID(userCryptoAddy.PrivateKey, chainID)
	if utility.FailOnError(err, "bind.NewKeyedTransactorWithChainID") == true {
		return false
	}

	auth.GasPrice = gasPrince
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := todo.DeployTodo(auth, con)
	if utility.FailOnError(err, "todo.DeployTodo") == true {
		return false
	}

	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

	return true
}
