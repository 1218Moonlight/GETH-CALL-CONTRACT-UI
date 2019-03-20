package Contract

import "github.com/ethereum/go-ethereum/common"

type commonAddress = common.Address

func HexToAddress(s string) commonAddress {
	return common.HexToAddress(s)
}
