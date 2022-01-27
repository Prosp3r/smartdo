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
	"github.com/holiman/uint256"
)

/*
PREVIOUS DEPLOYMENT
-------------------RESPONSE FROM TRANSACTION--------------------
CONTRACT HEX: 0x4241D10e086895Ca1E08903baB2778e49aa31d37
CREATION TRANSACTION HEX : 0x0cd557dd619d1f3e0c426fd67a8f710f8baa93ac03b2a3e36f577e53ac93b9c1
-------------------END RESPONSE FROM TRANSACTION----------------
*/

type InputQuery struct {
	AddressSpender   common.Address `json:"addressspender,omitempty"`
	Amount           uint256.Int    `json:"amount,omitempty"`
	SubtractedValue  uint256.Int    `json:"subtractedvalue,omitempty"`
	AddedValue       uint256.Int    `json:"addedvalue,omitempty"`
	AddressTo        common.Address `json:"addressto,omitempty"`
	AddressRecipient common.Address `json:"addressrecipient,omitempty"`
	AddressSender    common.Address `json:"addresssender,omitempty"`
	AddressNewOwner  common.Address `json:"addressnewowner,omitempty"`
	AddressOwner     common.Address `json:"addressowner,omitempty"`
	AddressAccount   common.Address `json:"addressaccount,omitempty"`
	Username         string         `json:"username"`
	Password         string         `json:"password"`
	CcontractHex     string         `json:"CcontractHex"`
}

type Response struct{}

var InputResponse Response

type PreTransactor struct {
	td  *todo.Todo
	bnd *bind.TransactOpts
	// addy *keystore.Key
}

func PreTransaction(con *ethclient.Client, input *InputQuery) (*PreTransactor, error) {

	var pt PreTransactor

	//Get user's crypto wallet
	userCryptoAddy, err := utility.ReadCryptoKey(input.Username, input.Password)
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

	contractAddy := common.HexToAddress(input.CcontractHex)

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

//
func Approve(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func DecreaseAllowance(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func IncreaseAllowance(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//Mint - Create more tokens
func Mint(ccon *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func RenounceOwnership(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func Transfer(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func TransferFrom(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func TransferOwnership(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func Allowance(con *ethclient.Client, input *InputQuery) (*uint256.Int, error) {
	//implement me
	return nil, nil
}

//
func BalanceOf(con *ethclient.Client, input *InputQuery) (*uint256.Int, error) {
	//implement me
	return nil, nil
}

//
func Decimal(con *ethclient.Client, input *InputQuery) (*uint8, error) {
	//implement me
	return nil, nil
}

//
func Name(con *ethclient.Client, input *InputQuery) (*string, error) {
	//implement me
	return nil, nil
}

//
func Owner(con *ethclient.Client, input *InputQuery) (*common.Address, error) {
	//implement me
	return nil, nil
}

//
func Symbol(con *ethclient.Client, input *InputQuery) (*string, error) {
	//implement me
	return nil, nil
}

//TotalSupply - Returns total amount of tokens in existence
func TotalSupply(con *ethclient.Client, input *InputQuery) (*uint256.Int, error) {
	//implement me
	pt, err := PreTransaction(con, input)
	if utility.FailOnError(err, "PreTransaction") == true {
		return nil, err
	}
	// pt.bnd
	// tdo := todo.NewTodo(input.CcontractHex)
	fmt.Println(pt)

	return nil, nil
}

// func InteractAdd(con *ethclient.Client, username, password, CcontractHex string) bool {

// 	preTrans, err := PreTransaction(con, username, password, CcontractHex)
// 	if utility.FailOnError(err, "PreTransaction") == true {
// 		return false
// 	}

// 	transaction, err := preTrans.td.Add(preTrans.bnd, "MAKE BURGER")
// 	if utility.FailOnError(err, "it.Add") == true {
// 		return false
// 	}

// 	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
// 	fmt.Printf("Transaction Hash: %v\n", transaction.Hash())
// 	fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
// 	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

// 	return true
// }

// func InteractList(con *ethclient.Client, username, password, CcontractHex string) (*[]todo.TodoTask, error) {

// 	preTrans, err := PreTransaction(con, username, password, CcontractHex)
// 	if utility.FailOnError(err, "PreTransaction") == true {
// 		return nil, err
// 	}

// 	transaction, err := preTrans.td.List(&bind.CallOpts{
// 		From: preTrans.bnd.From,
// 	})
// 	if utility.FailOnError(err, "preTrans.td.List") == true {
// 		return nil, err
// 	}

// 	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
// 	fmt.Printf("Transaction : %v\n", transaction)
// 	// fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
// 	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

// 	return &transaction, nil
// }

// func InteractUpdate(con *ethclient.Client, username, password, CcontractHex, ItemName, updatedTask string) bool {

// 	preTrans, err := PreTransaction(con, username, password, CcontractHex)
// 	if utility.FailOnError(err, "PreTransaction") == true {
// 		return false
// 	}

// 	var itemId *big.Int
// 	iid, _ := InteractList(con, username, password, CcontractHex)
// 	for i, v := range *iid {
// 		if v.Content == ItemName {
// 			itemId = big.NewInt(int64(i))
// 		}
// 	}
// 	transaction, err := preTrans.td.Update(preTrans.bnd, itemId, updatedTask)
// 	if utility.FailOnError(err, "ipreTrans.td.Update") == true {
// 		return false
// 	}

// 	fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
// 	fmt.Printf("Transaction Hash: %v\n", transaction.Hash())
// 	fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
// 	fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

// 	return true
// }

// func InteractRemove(con *ethclient.Client, username, password, CcontractHex string) bool {

// 	// preTrans, err := PreTransaction(con, username, password, CcontractHex)
// 	// if utility.FailOnError(err, "PreTransaction") == true {
// 	// 	return false
// 	// }

// 	// transaction, err := preTrans.td.Add(preTrans.bnd, "MAKE BURGER")
// 	// if utility.FailOnError(err, "it.Add") == true {
// 	// 	return false
// 	// }

// 	// fmt.Println("-------------------RESPONSE FROM TRANSACTION--------------------")
// 	// fmt.Printf("Transaction Hash: %v\n", transaction.Hash())
// 	// fmt.Printf("Transaction Hex: %v\n", transaction.Hash().Hex())
// 	// fmt.Println("-------------------END RESPONSE FROM TRANSACTION----------------")

// 	return true
// }
