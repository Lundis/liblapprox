package discrete

import (
    "math"
    "code.google.com/p/liblundis/lmath/util/cont"
)


func Max(values []float64) (max float64, index int) {
    max = values[0]
    index = 0
    for i, v := range values {
        if v > max {
            max = v
            index = i
        }
    }
    return max, index
}

func Min(values []float64) (min float64, index int) {
    min = values[0]
    index = 0
    for i, v := range values {
        if v < min {
            min = v
            index = i
        }
    }
    return min, index
}

func VectorAbsDiff(x, y []float64) []float64 {
    z := make([]float64, len(x))
    for i := range x {
        z[i] = math.Abs(x[i] - y[i])
    }
    return z
}

func Values(f cont.Function, x []float64) []float64 {
    y := make([]float64, len(x))
    for i, v := range x {
        y[i] = f(v)
    }
    return y
}

func FindMaxDiff(x, y []float64) (max float64, index int) {
    diff := VectorAbsDiff(x, y)
    return Max(diff)
}

func Minus(x, y []float64) []float64 {
    z := make([]float64, len(x))
    for i := range z {
        z[i] = x[i] - y[i]
    }
    return z
}

func Plus(x, y []float64) []float64 {
    z := make([]float64, len(x))
    for i := range z {
        z[i] = x[i] + y[i]
    }
    return z
}