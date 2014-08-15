package liblundis

import "math"

func Equals(x, y float64) bool {
    return math.Abs(x-y) < 1e-6
}

func MapKeysIntToSlice(m map[int]interface{}) []int {
	s := make([]int, 0, len(m))
	for key := range m {
		s = append(s, key)
	}
	return s
}