package Contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

type abiABI = abi.ABI

const (
	balanceAbi = `[{"constant": true,"inputs": [{"name": "_owner","type": "address"}],"name": "balanceOf", "outputs": [{"name": "balance","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}]`
	symbolAbi  = `[{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name": "","type": "string"}],"payable":false,"stateMutability":"view","type":"function"}]`
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
