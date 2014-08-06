package exchange

import (
	"code.google.com/p/liblundis/lmath"
)

type Iteration struct {
	Error float64
	Poly Polynomial
}

// Approximates for the specified degrees, saving the result in approx.
// Returns details about the iterations.
func Approximate(approx *approx.Approx, degrees []int) map[int] []Iteration {
	approx.
}

// Approximates one degree, saving the result in approx
// Returns information about the iterations
func ApproximateDegree(f lmath.Func1to1, degree int, approx *approx.Approx) []Iteration {
	roots := ipol.GenerateChebyshevRoots(degree+2, approx.Start, approx.End)

}

func SolveEquationSystem(approx approx.Approx, roots []float64) *Iteration {
	degree := len(roots)-2
	matrix := make([][]float64, len(roots))
	for row := range matrix {
		x := roots[row]
		// degree n has n+1 columns
		// + one col for error, one for f[roots[col]]
		equation := make([]float64, degree + 3)
		for col = 0; col <= degree; col++ {
			equation[col] = math.Pow(x, col)
		}
		// error with alternating signs
		equation[degree+1] = math.Pow(-1, row % 2)
		equation[degree+2] = approx.Func(x)
		matrix[row] = equation
	}
	
}