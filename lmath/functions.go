package lmath

import (
    "math"
    //"fmt"
)

type Func1to1 func(float64) float64

type Vector []float64


func ZeroFunc() Func1to1 {
    return func(x float64) float64 {
        return 0
    }
}

func Max(f Func1to1, accuracy int, start, end float64) float64 {
    max := float64(0)
    for i := 0; i <= accuracy; i++ {
        max = math.Max(max, f(start + (end - start) * (float64(i)/float64(accuracy))))
    }
    return max
}

// Calculates a function repeatedly with increasing accuracy until it converges (less than 1e-6 difference).
// Warning: can be very inefficient, crash and/or loop indefinitely for "advanced" functions
func CalculateAccurately(f func(accuracy int) float64, initial_accuracy int) float64 {
    accuracy := initial_accuracy
    iter1 := float64(0)
    iter2 := f(accuracy)
    for math.Abs(iter1 - iter2) > 1e-6 {
        accuracy *= 2
        iter1 = iter2
        iter2 = f(accuracy)
    }
    return iter2
}

func FuncAbsDiff(f, g Func1to1) Func1to1 {
    return func(x float64) float64 {
        return math.Abs(f(x) - g(x))
    }
}

func Values(f Func1to1, x Vector) Vector {
    y := make([]float64, len(x))
    for i, v := range x {
        y[i] = f(v)
    }
    return y
}

func FindMaxDiff(f, g Func1to1, start, end float64) float64 {
    fgdiff := FuncAbsDiff(f, g)
    error := CalculateAccurately(func(accuracy int) float64 {
        return Max(fgdiff, accuracy, start, end) 
        }, 100)
    return error
}

