package liblundis

import "math"

func Equals(x, y float64) bool {
    return math.Abs(x-y) < 1e-6
}