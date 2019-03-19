package Contract

import (
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	geth "github.com/ethereum/go-ethereum/ethclient"
	common "github.com/ethereum/go-ethereum/common"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
)

func GetAbi(stringAbi string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(stringAbi))
}

func Dial(url string) (*geth.Client, error) { // "http://[ip]:[port]"
	return geth.Dial(url)
}

func NewBoundContract(address common.Address, abi abi.ABI, caller bind.ContractCaller)(*bind.BoundContract){
	return &bind.BoundContract{address, abi, caller, nil, nil}
}

func HexToAddress(s string)common.Address{
	return common.HexToAddress(s)
}
