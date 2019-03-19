package Contract

import (
	"testing"
)

func TestContract(t *testing.T) {
	_, err := GetBalanceAbi()
	if err != nil {
		t.Error(err)
		return
	}
}
