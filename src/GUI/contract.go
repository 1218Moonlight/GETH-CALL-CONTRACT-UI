package GUI

import (
	contract "github.com/Kioryu/GETH-CALL-CONTRACT-UI/Contract"
	"math/big"
)

const host = "http://192.168.0.133:8545"

func symbol(contractAddress string) (*string, error) {
	symbol, abi, err := contract.GetSymbolAbi()
	if err != nil {
		return nil, err
	}

	client, err := contract.Dial(host)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ca := contract.NewBoundContract(contract.HexToAddress(contractAddress), abi, client)

	err = ca.Call(nil, symbol, "symbol")
	if err != nil {
		return nil, err
	}

	return symbol, nil
}

func balanceOf(contractAddress string, externalOwnedAccount string) (*big.Int, error) {
	balance, abi, err := contract.GetBalanceOfAbi()
	if err != nil {
		return nil, err
	}

	client, err := contract.Dial(host)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ca := contract.NewBoundContract(contract.HexToAddress(contractAddress), abi, client)
	err = ca.Call(nil, balance, "balanceOf", contract.HexToAddress(externalOwnedAccount))
	if err != nil {
		return nil, err
	}

	return *balance, nil
}
