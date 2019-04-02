package Contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"math/big"
)

type (
	abiABI = abi.ABI
)

const (
	nameAbi     = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`
	totalSupply = `[{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	balanceAbi  = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name": "balance","type": "uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	symbolAbi   = `[{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`
	decimals    = `[{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"}]`
	allowance   = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	erc20Abi    = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name": "balance","type": "uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
)

type Erc20Abis struct {
	balanceOf **big.Int
	symbol    *string
}

func GetAbi(stringAbi string) (abiABI, error) {
	return abi.JSON(strings.NewReader(stringAbi))
}

func GetErc20Abi() (Erc20Abis, abiABI, error) {
	eAbi, err := GetAbi(erc20Abi)
	return Erc20Abis{balanceOf: new(*big.Int), symbol: new(string)}, eAbi, err
}

func GetBalanceOfAbi() (**big.Int, abiABI, error) {
	bAbi, err := GetAbi(balanceAbi)
	return new(*big.Int), bAbi, err
}

func GetSymbolAbi() (*string, abiABI, error) {
	sAbi, err := GetAbi(symbolAbi)
	return new(string), sAbi, err
}

func GetNameAbi() (*string, abiABI, error) {
	nAbi, err := GetAbi(nameAbi)
	return new(string), nAbi, err
}

func GetTotalSupplyAbi() (**big.Int, abiABI, error) {
	tAbi, err := GetAbi(totalSupply)
	return new(*big.Int), tAbi, err
}

func GetDecimalsAbi() (**big.Int, abiABI, error) {
	dAbi, err := GetAbi(decimals)
	return new(*big.Int), dAbi, err
}

func GetAllowance() (**big.Int, abiABI, error) {
	aAbi, err := GetAbi(allowance)
	return new(*big.Int), aAbi, err
}
