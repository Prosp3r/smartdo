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

type PreTransactor struct {
	td  *todo.Todo
	bnd *bind.TransactOpts
	// addy *keystore.Key
}

func PreTransaction(con *ethclient.Client, username, password, CcontractHex string) (*PreTransactor, error) {

	var pt PreTransactor

	//Get user's crypto wallet
	userCryptoAddy, err := utility.ReadCryptoKey(password, username)
	if utility.FailOnError(err, "utility.ReadCryptoKey") == true {
		return nil, err
	}

	nonce, err := con.PendingNonceAt(context.Background(), userCryptoAddy.Address)
	if utility.FailOnError(err, "con.PendingNonceAt") == true {
		return nil, err
	}

	gasPrince, err := con.SuggestGasPrice(context.Background())
	if utility.FailOnError(err, "con.SuggestGasPrice") == true {
		return nil, err
	}

	chainID, err := con.NetworkID(context.Background())
	if utility.FailOnError(err, "con.NetworkID") == true {
		return nil, err
	}

	contractAddy := common.HexToAddress(CcontractHex)

	it, err := todo.NewTodo(contractAddy, con)
	if utility.FailOnError(err, "todo.NewTodo") == true {
		return nil, err
	}

	tx, err := bind.NewKeyedTransactorWithChainID(userCryptoAddy.PrivateKey, chainID)
	if utility.FailOnError(err, "bind.NewKeyedTransactorWithChainID") == true {
		return nil, err
	}

	tx.GasPrice = gasPrince
	tx.GasLimit = uint64(3000000)
	tx.Nonce = big.NewInt(int64(nonce))

	pt.bnd = tx
	pt.td = it
	// pt.addy = userCryptoAddy

	return &pt, nil
}

func InteractAdd(con *ethclient.Client, username, password, CcontractHex string) bool {

	preTrans, err := PreTransaction(con, username, password, CcontractHex)
	if utility.FailOnError(err, "PreTransaction") == true {
		return false
	}

	transaction, err := preTrans.td.Add(preTrans.bnd, "MAKE BURGER")
	if utility.FailOnError(err, "it.Add") == true {
		return false
	}

	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
	fmt.Printf("Transaction Hash: %v\n", transaction.Hash())
	fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

	return true
}

func InteractList(con *ethclient.Client, username, password, CcontractHex string) bool {

	preTrans, err := PreTransaction(con, username, password, CcontractHex)
	if utility.FailOnError(err, "PreTransaction") == true {
		return false
	}

	transaction, err := preTrans.td.List(&bind.CallOpts{
		From: preTrans.bnd.From,
	})
	if utility.FailOnError(err, "preTrans.td.List") == true {
		return false
	}

	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
	fmt.Printf("Transaction : %v\n", transaction)
	// fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

	return true
}

func InteractRemove(con *ethclient.Client, username, password, CcontractHex string) bool {

	// preTrans, err := PreTransaction(con, username, password, CcontractHex)
	// if utility.FailOnError(err, "PreTransaction") == true {
	// 	return false
	// }

	// transaction, err := preTrans.td.Add(preTrans.bnd, "MAKE BURGER")
	// if utility.FailOnError(err, "it.Add") == true {
	// 	return false
	// }

	// fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
	// fmt.Printf("Transaction Hash: %v\n", transaction.Hash())
	// fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
	// fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

	return true
}
