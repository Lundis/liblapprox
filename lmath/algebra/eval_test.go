package algebra

import(
	"testing"
	"math"
	"code.google.com/p/liblundis/lmath"
)

func TestEvalCommonf64(t *testing.T) {
	tests := map[string] float64 {
		"1 + 2 + 3": 6,
		"pi + 1": math.Pi + 1,
		"e + 1": math.E + 1,
		"e / pi": math.E / math.Pi,
		"(e + 2*pi) / pi": (math.E + 2*math.Pi) / math.Pi,
	}
	for str, expected:= range tests {
		result, err := EvalCommonf64(str)
		if err != nil {
			t.Error(err.Error())
		} else {
			lmath.AssertEqualsFloat64(t, result, expected, "Error")
		}
	}
}