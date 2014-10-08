package cont

import (
    "math"
    //"fmt"
)

type Function func(float64) float64

type Vector []float64


func ZeroFunc() Function {
    return func(x float64) float64 {
        return 0
    }
}

func Max(f Function, accuracy int, start, end float64) (max float64, loc float64) {
    max = f(start)
    loc = float64(start)
    for i := 0; i < accuracy; i++ {
        x := start + (end - start) * (float64(i)/float64(accuracy-1))
        y := f(x)
        if y > max {
            max = y
            loc = x
        }
    }
    return max, loc
}

func Min(f Function, accuracy int, start, end float64) (min float64, loc float64) {
    min = f(start)
    loc = float64(start)
    for i := 0; i < accuracy; i++ {
        x := start + (end - start) * (float64(i)/float64(accuracy-1))
        y := f(x)
        if y < min {
            min = y
            loc = x
        }
    }
    return min, loc
}

// Calculates a function repeatedly with increasing accuracy until it converges (less than 1e-6 difference).
// Warning: can be very inefficient, crash and/or loop indefinitely for "advanced" functions
func CalculateAccurately(f func(accuracy int) (float64, float64), initial_accuracy int) (float64, float64) {
    accuracy := initial_accuracy
    iter1 := float64(0)
    iter2, loc := f(accuracy)
    for math.Abs(iter1 - iter2) > 1e-6 && accuracy < initial_accuracy*2^8 {
        accuracy *= 2
        iter1 = iter2
        iter2, loc = f(accuracy)
    }
    return iter2, loc
}

func FuncAbsDiff(f, g Function) Function {
    return func(x float64) float64 {
        return math.Abs(f(x) - g(x))
    }
}

func FindMaxDiff(f, g Function, start, end float64) (max_diff, loc float64) {
    fgdiff := FuncAbsDiff(f, g)
    max_diff, loc = CalculateAccurately(func(accuracy int) (float64, float64) {
        return Max(fgdiff, accuracy, start, end)
    }, 128)
    return max_diff, loc
}

// takes a monotonically increasing function and finds the location where it's within accuracy from value
func BinarySearch(f Function, value, start, end float64) float64 {
    low_loc := start
    high_loc := end
    for loc := (start + end) / 2;; loc = (high_loc + low_loc) / 2 {
        y := f(loc)
        if y == value || loc == low_loc || loc == high_loc {
            return loc
        } else if y > value { // we need to go lower
            high_loc = loc
        } else {
            low_loc = loc
        }
    }
}