package GUI

import (
	contract "github.com/Kioryu/GETH-CALL-CONTRACT-UI/Contract"
	"math/big"
)

const host = "http://192.168.0.133:8545"

func erc20All(contractAddress string, eoa string) (contract.Erc20Abis, error) {
	erc20struct, abi, err := contract.GetErc20Abi()
	if err != nil {
		return contract.Erc20Abis{}, err
	}

	client, err := contract.Dial(host)
	if err != nil {
		return contract.Erc20Abis{}, err
	}
	defer client.Close()

	ca := contract.NewBoundContract(contract.HexToAddress(contractAddress), abi, client)

	err = ca.Call(nil, erc20struct.BalanceOf, "balanceOf", contract.HexToAddress(eoa))
	if err != nil {
		return contract.Erc20Abis{}, err
	}

	err = ca.Call(nil, erc20struct.Symbol, "symbol")
	if err != nil {
		return contract.Erc20Abis{}, err
	}

	err = ca.Call(nil, erc20struct.Name, "name")
	if err != nil {
		return contract.Erc20Abis{}, err
	}

	err = ca.Call(nil, erc20struct.Decimals, "decimals")
	if err != nil {
		return contract.Erc20Abis{}, err
	}

	err = ca.Call(nil, erc20struct.TotalSupply, "totalSupply")
	if err != nil {
		return contract.Erc20Abis{}, err
	}

	return erc20struct, nil
}

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
