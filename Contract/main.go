package Contract

import (
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

const (
	balanceAbi = `[{"constant": true,"inputs": [{"name": "_owner","type": "address"}],"name": "balanceOf", "outputs": [{"name": "balance","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}]`
)

func GetBalanceAbi() (abi.ABI, error) {
	balanceOfAbi, err := abi.JSON(strings.NewReader(balanceAbi))
	if err != nil {
		return abi.ABI{}, err
	}
	return balanceOfAbi, nil
}
