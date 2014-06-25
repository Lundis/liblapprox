package lmath

// http://en.wikipedia.org/wiki/Binomial_coefficient
func BinCoeff(n, k int) int64 {
	product := int64(1)
	for i := int64(k + 1); i <= int64(n); i++ {
		product *= i
	}
	for i := int64(2); i <= int64(n - k); i++ {
		product /= i
	}
	return product
}