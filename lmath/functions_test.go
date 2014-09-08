package lmath

import(
	"math"
	"math/rand"
	"testing"
)

func TestMax(t *testing.T) {
    
    for i := 0; i < 10; i++ {
    	x0 := float64(rand.Float32())
    	f := func(x float64) float64 {
    		return -math.Pow(x-x0, 2.0)
    	}
    	max, loc := Max(f, 400, x0 - 1.054213, x0 + 1.1236578234)
    	if !EqualsFloat(max, 0, 1e-6) {
    		t.Errorf("max == %v. Should be 0", max)
    	}
        if math.Abs(loc - x0) > 0.001 {
            t.Errorf("loc == %v. Should be %v", loc, x0)
        }
    }
}

func TestCalculateAccurately(t *testing.T) {

}

func TestBinarySearch(t *testing.T) {
    f := func(x float64) float64 {
        return x*x
    }

    loc := BinarySearch(f, 1, 0, 2)
    if loc != 1 {
        t.Errorf("got %v, expected 1", loc)
    }

    // make sure that it converges at the end of the interval
    loc = BinarySearch(f, 2, 0, 1)
    if loc != 1 {
        t.Errorf("got %v, expected 1", loc)
    }

}