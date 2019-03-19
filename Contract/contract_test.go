package Contract

import (
	"testing"
	"math/big"
)

const (
	balanceAbi = `[{"constant": true,"inputs": [{"name": "_owner","type": "address"}],"name": "balanceOf", "outputs": [{"name": "balance","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}]`
	ca = "0x1d4c85d521ca26177772639ee00609eb573f043c"
	eoa = "0x6f090f6cb125f77396d4b8f52fdabf7d5c1b53d4"
)

func TestBalanceOf(t *testing.T) {
	abi, err := GetAbi(balanceAbi)
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

	contract := NewBoundContract(HexToAddress(ca), abi, client)

	balance := new(*big.Int)
	err = contract.Call(nil, balance, "balanceOf", HexToAddress(eoa))
	if err != nil {
		t.Error(err)
		return
	}

}
