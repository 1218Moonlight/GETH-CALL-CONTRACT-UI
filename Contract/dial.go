package Contract

import geth "github.com/ethereum/go-ethereum/ethclient"

func Dial(url string) (*geth.Client, error) { // "http://[ip]:[port]"
	return geth.Dial(url)
}
