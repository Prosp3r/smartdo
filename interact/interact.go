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
func Mint(con *ethclient.Client, ii *InputQuery) bool {
	//implement me
	pt, err := PreTransaction(con, ii)
	if utility.FailOnError(err, "interact.PreTransaction") == true {
		fmt.Println("Cound not prepare pretransactor - ", err)
		return false
	}

	contractAdd := common.HexToAddress(ii.CcontractHex)
	cTodo, err := todo.NewTodo(contractAdd, con)
	if utility.FailOnError(err, "Creating contract bindings") == true {
		fmt.Println("Could not access contract bindings")
	}

	cTodo.Mint(pt.bnd, ii.AddressRecipient, ii.Amount.ToBig())
	return false
}

//
func RenounceOwnership(con *ethclient.Client, input *InputQuery) (*bool, error) {
	//implement me
	return nil, nil
}

//
func Transfer(con *ethclient.Client, ii *InputQuery) bool {
	//implement me
	pt, err := PreTransaction(con, ii)
	if utility.FailOnError(err, "interact.PreTransaction") == true {
		fmt.Println("Cound not prepare pretransactor - ", err)
		return false
	}

	contractAdd := common.HexToAddress(ii.CcontractHex)
	cTodo, err := todo.NewTodo(contractAdd, con)
	if utility.FailOnError(err, "Creating contract bindings") == true {
		fmt.Println("Could not access contract bindings")
	}

	cTodo.Transfer(pt.bnd, ii.AddressRecipient, ii.Amount.ToBig())
	return false
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

	return nil, nil
}
