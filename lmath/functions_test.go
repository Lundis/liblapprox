package lmath

import(
	"math"
	"math/rand"
	"testing"
    "code.google.com/p/liblundis"
)

func TestMax(t *testing.T) {
    
    for i := 0; i < 10; i++ {
    	x0 := float64(rand.Float32())
    	f := func(x float64) float64 {
    		return -math.Pow(x-x0, 2.0)
    	}
    	max := Max(f, 200, x0 - 1.054213, x0 + 1.1236578234)
    	if !liblundis.Equals(max, 0) {
    		t.Errorf("max == %v. Should be 0", max)
    	}
    }
}

func TestCalculateAccurately(t *testing.T) {

}