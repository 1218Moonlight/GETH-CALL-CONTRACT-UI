package Contract

import (
	"testing"
)

const (
	ca  = "0x55d8Fe2df965043486662A21A16E37DA4E52981f"
	eoa = "0x6f090f6cb125f77396d4b8f52fdabf7d5c1b53d4"
)

func TestErc20Abi(t *testing.T){
	erc20struct, abi, err := GetErc20Abi()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := Dial("http://192.168.0.133:8545")
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Close()

	ca := NewBoundContract(HexToAddress(ca), abi, client)
	err = ca.Call(nil, erc20struct.BalanceOf, "balanceOf", HexToAddress(eoa))
	if err != nil {
		t.Error(err)
		return
	}

	err = ca.Call(nil, erc20struct.Symbol, "symbol")
	if err != nil {
		t.Error(err)
		return
	}

	err = ca.Call(nil, &erc20struct.Name, "name")
	if err != nil {
		t.Error(err)
		return
	}


}

func TestName(t *testing.T) {
	name, abi, err := GetNameAbi()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := Dial("http://192.168.0.133:8545")
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Close()

	ca := NewBoundContract(HexToAddress(ca), abi, client)

	err = ca.Call(nil, name, "name")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBalanceOf(t *testing.T) {
	balance, abi, err := GetBalanceOfAbi()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := Dial("http://192.168.0.133:8545")
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Close()

	ca := NewBoundContract(HexToAddress(ca), abi, client)
	err = ca.Call(nil, balance, "balanceOf", HexToAddress(eoa))
	if err != nil {
		t.Error(err)
		return
	}

}

func TestSymbol(t *testing.T) {
	symbol, abi, err := GetSymbolAbi()
	if err != nil {
		t.Error(err)
		return
	}

	client, err := Dial("http://192.168.0.133:8545")
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Close()

	ca := NewBoundContract(HexToAddress(ca), abi, client)

	err = ca.Call(nil, symbol, "symbol")
	if err != nil {
		t.Error(err)
		return
	}
}
