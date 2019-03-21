package Contract

import (
	"testing"
)

const (
	ca  = "0x1d4c85d521ca26177772639ee00609eb573f043c"
	eoa = "0x6f090f6cb125f77396d4b8f52fdabf7d5c1b53d4"
)

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
