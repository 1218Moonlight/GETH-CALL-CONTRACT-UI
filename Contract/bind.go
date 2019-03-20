package Contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewBoundContract(address commonAddress, abi abiABI, caller bind.ContractCaller) *bind.BoundContract {
	return bind.NewBoundContract(address, abi, caller, nil, nil)
}
