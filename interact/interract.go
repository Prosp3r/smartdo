package interact

import (
	"context"
	"fmt"
	"math/big"

	todo "github.com/Prosp3r/smartdo/gen"
	"github.com/Prosp3r/smartdo/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
PREVIOUS DEPLOYMENT
-------------------RESPONSE FROM TRANSACTION--------------------
CONTRACT HEX: 0x4241D10e086895Ca1E08903baB2778e49aa31d37
CREATION TRANSACTION HEX : 0x0cd557dd619d1f3e0c426fd67a8f710f8baa93ac03b2a3e36f577e53ac93b9c1
-------------------END RESPONSE FROM TRANSACTION----------------
*/

func Interact(con *ethclient.Client, username, password, CcontractHex string) bool {
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

	contractAddy := common.HexToAddress(CcontractHex)

	it, err := todo.NewTodo(contractAddy, con)
	if utility.FailOnError(err, "todo.NewTodo") == true {
		return false
	}

	tx, err := bind.NewKeyedTransactorWithChainID(userCryptoAddy.PrivateKey, chainID)
	if utility.FailOnError(err, "bind.NewKeyedTransactorWithChainID") == true {
		return false
	}

	tx.GasPrice = gasPrince
	tx.GasLimit = uint64(3000000)
	tx.Nonce = big.NewInt(int64(nonce))

	transaction, err := it.Add(tx, "MAKE BURGER")
	if utility.FailOnError(err, "it.Add") == true {
		return false
	}

	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
	fmt.Printf("Transaction Hash: %v\n", transaction.Hash())
	fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

	return true
}
