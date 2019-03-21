package Contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

type abiABI = abi.ABI

const (
	balanceAbi  = `[{"constant":true,"inputs":[{"name": "_owner","type": "address"}],"name": "balanceOf", "outputs": [{"name": "balance","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}]`
	symbolAbi   = `[{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name": "","type": "string"}],"payable":false,"stateMutability":"view","type":"function"}]`
	nameAbi     = `{"constant": true,"inputs": [],"name": "name","outputs": [{"name": "","type": "string"}],"payable": false,"stateMutability": "view","type": "function"}`
	totalSupply = `{"constant": true,"inputs": [],"name": "totalSupply","outputs": [{"name": "","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}`
	decimals    = `{"constant": true,"inputs": [],"name": "decimals","outputs": [{"name": "","type": "uint8"}],"payable": false,"stateMutability": "view","type": "function"}`
	allowance   = `{"constant": true,"inputs": [{"name": "_owner","type": "address"},{"name": "_spender","type": "address"}],"name": "allowance","outputs": [{"name": "","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}`
)

func GetAbi(stringAbi string) (abiABI, error) {
	return abi.JSON(strings.NewReader(stringAbi))
}

func GetBalanceOfAbi() (abiABI, error) {
	return GetAbi(balanceAbi)
}

func GetSymbolAbi() (abiABI, error) {
	return GetAbi(symbolAbi)
}

func GetNameAbi() (abiABI, error) {
	return GetAbi(nameAbi)
}

func GetTotalSupplyAbi() (abiABI, error) {
	return GetAbi(totalSupply)
}

func GetDecimalsAbi() (abiABI, error) {
	return GetAbi(decimals)
}

func GetAllowance() (abiABI, error) {
	return GetAbi(allowance)
}
