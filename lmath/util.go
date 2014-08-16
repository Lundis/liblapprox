package lmath

import (
	"math/big"
	"math"
)

func EqualsFloat(x, y, accuracy float64) bool {
    return math.Abs(x-y) < accuracy
}

func MapKeysIntToSlice(m map[int]interface{}) []int {
	s := make([]int, 0, len(m))
	for key := range m {
		s = append(s, key)
	}
	return s
}

func NewRatf(k float64) *big.Rat {
	return big.NewRat(1,1).SetFloat64(k)
}

func NewRati(k int) *big.Rat {
	return big.NewRat(int64(k),1)
}

func NewRati64(k int64) *big.Rat {
	return big.NewRat(k,1)
}

func NewRat(r *big.Rat) *big.Rat {
	return big.NewRat(1,1).Set(r)
}

// Returns a big.Rat
func BigRatPow(r *big.Rat, k int) (dst *big.Rat) {
	dst = big.NewRat(1,1)
	for i := 0; i < k; i++ {
		dst.Mul(dst, r)
	}
	return dst
}