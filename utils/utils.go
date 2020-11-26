package utils

import (
	"fmt"
	"math"
	"math/big"
)

// WeiToEther convert wei to ether.
func WeiToEther(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance,big.NewFloat(math.Pow10(18)))
	fmt.Printf("Ether: %v\n",ethValue)
	return ethValue
}

