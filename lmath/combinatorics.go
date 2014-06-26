package lmath

import "math/big"

// http://en.wikipedia.org/wiki/Binomial_coefficient
func BinCoeff(n, k int) int64 {
	product := big.NewInt(1)
	return product.Binomial(int64(n), int64(k)).Int64()
}