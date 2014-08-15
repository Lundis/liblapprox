package exchange

import (
	"code.google.com/p/liblundis/lmath"
	"code.google.com/p/liblundis/lmath/algebra"
	"code.google.com/p/liblundis/lmath/ipol"
	"code.google.com/p/liblundis/lmath/approx"
	"math"
	"fmt"
)

type Iteration struct {
	Max_error, Leveled_error float64
	Poly lmath.Polynomial
}

func (self Iteration) String() string {
	format := "%v || error: %.6f || error_diff: %.6f"
	return fmt.Sprintf(format, self.Poly, self.Max_error, self.ErrorDiff)
}

func (self Iteration) ErrorDiff() float64 {
	return self.Max_error - self.Leveled_error
}

// Approximates for the specified degrees, saving the result in approx.
// Returns details about the iterations.
func Approximate(approx *approx.Approx, degrees []int, accuracy float64) map[int] []Iteration {
	if accuracy <= 0 {
		panic("exchange.Approximate(): accuracy must be >0")
	}
	iters := make(map[int] []Iteration)
	for _, deg := range degrees {
		iters[deg] = ApproximateDegree(approx, deg, accuracy)
	}
	return iters
}

// Approximates one degree, saving the result in approx
// Returns information about the iterations
func ApproximateDegree(approx *approx.Approx, degree int, accuracy float64) []Iteration {
	iters := make([]Iteration, 0, 10)
	roots := ipol.GenerateChebyshevRoots(degree+2, approx.Start, approx.End)
	matrix := createMatrix(approx, roots)
	// random init value.
	diff := accuracy*2
	// TODO: remove iter limit when this actually converges
	for i := 0; diff > accuracy && i < 20; i++ {
		fmt.Printf("iter %v\n", i)
		updateMatrix(matrix, approx, roots)
		matrix.Solve()
		iter := Iteration{}
		iter.Poly, iter.Leveled_error = interpretSolution(matrix)
		var loc float64
		iter.Max_error, loc = lmath.FindMaxDiff(approx.Func, iter.Poly.Function(), approx.Start, approx.End)
		diff = iter.Max_error - iter.Leveled_error
		iters = append(iters, iter)
		updateRoots(roots, approx.Func, iter.Poly.Function(), loc)
	}
	return iters
}

func createMatrix(approx *approx.Approx, roots []float64) algebra.Matrix {
	rows := len(roots)
	cols := rows + 1
	return algebra.NewMatrix(rows, cols)
}

func updateMatrix(m algebra.Matrix, approx *approx.Approx, roots []float64) {
	degree := len(roots)-2
	for row, equation := range m {
		x := roots[row]
		for col := 0; col <= degree; col++ {
			equation[col] = math.Pow(x, float64(col))
		}
		// error with alternating signs
		equation[degree+1] = math.Pow(-1, float64(row % 2))
		// actual function value
		equation[degree+2] = approx.Func(x)
	}
}


// interprets the solution and returns the polynomial and the leveled error
func interpretSolution(matrix algebra.Matrix) (lmath.Polynomial, float64) {
	poly := lmath.NewPolynomial(len(matrix)-2)
	last_col := len(matrix[0]) - 1
	for row := 0; row < len(matrix)-1; row++ {
		poly[row] = matrix[row][last_col]
	}
	return poly, math.Abs(matrix[len(matrix)-1][last_col])
}

func updateRoots(roots []float64, orig_func, approx_func lmath.Func1to1, loc float64) {
	// find the first root larger than loc
	i := 0
	for ; i < len(roots)-1; i++ {
		if roots[i] >= loc {
			break;
		}
	}
	fmt.Printf("roots: %v", roots)
	fmt.Printf("max at %.4f\n", loc)
	// then replace either it or the previous/next one, depending on sign of errors
	root_error := orig_func(roots[i]) - approx_func(roots[i])
	max_error := orig_func(loc) - approx_func(loc)
	if math.Signbit(root_error) == math.Signbit(max_error) {
		fmt.Printf("replacing %.4f\n", roots[i])
		roots[i] = loc

	} else {
		if i == 0 {
			fmt.Printf("replacing %.4f\n", roots[i+1])
			roots[i+1] = loc
			// debug
			fmt.Printf("wat, this shouldnt happen\n")
		} else {
			fmt.Printf("replacing %.4f\n", roots[i-1])
			roots[i-1] = loc
		}
	}
}