package exchange

import (
	. "code.google.com/p/liblundis/lmath"
	"code.google.com/p/liblundis/lmath/base"
	"code.google.com/p/liblundis/lmath/algebra"
	"code.google.com/p/liblundis/lmath/ipol"
	"code.google.com/p/liblundis/lmath/approx"
	"math"
	"fmt"
	"sort"
)

type Iteration struct {
	Max_error, Leveled_error float64
	Basis base.Basis
	Replacing float64
	New_root float64
}

func (self Iteration) String() string {
	format := "%v || error: %.6f || error_diff: %.6f"
	return fmt.Sprintf(format, self.Basis, self.Max_error, self.ErrorDiff())
}

func (self Iteration) ErrorDiff() float64 {
	return self.Max_error - self.Leveled_error
}

// Approximates for the specified degrees, saving the result in approx.
// Returns details about the iterations.
func Approximate(approx *approx.Approx, degrees []int, accuracy float64, basis_func base.BasisImplConverter) map[int] []Iteration {
	if accuracy <= 0 {
		panic("exchange.Approximate(): accuracy must be >0")
	}
	iters := make(map[int] []Iteration)
	for _, deg := range degrees {
		iters[deg] = ApproximateDegree(approx, deg, accuracy, basis_func)
	}
	return iters
}

// Approximates one degree, saving the result in approx
// Returns information about the iterations
func ApproximateDegree(approx *approx.Approx, degree int, accuracy float64, basis_func base.BasisImplConverter) []Iteration {
	iters := make([]Iteration, 0, 10)
	roots := ipol.GenerateChebyshevRoots(degree+2, approx.Start, approx.End)
	matrix := createMatrix(approx, roots)
	quit := false
	old_leveled_error := float64(0)
	for i := 0; !quit && i < 100; i++ {
		updateMatrix(matrix, approx, roots)
		matrix.Solve()
		iter := Iteration{}
		var b *base.BasisImpl
		b, iter.Leveled_error = interpretSolution(matrix)
		iter.Basis = basis_func(b)
		var loc float64
		iter.Max_error, loc = FindMaxDiff(approx.Func, iter.Basis.Function(), approx.Start, approx.End)
		diff := iter.Max_error - iter.Leveled_error
		if old_leveled_error >= iter.Leveled_error {
			fmt.Printf("exchange failed to converge for deg %v!\n", degree)
			// if the new leveled error is smaller than the old the algorithm is broken and should stop
			quit = true
		} else if diff < accuracy {
			quit = true
			old_leveled_error = iter.Leveled_error
		}
		iter.Replacing = updateRoots(roots, approx.Func, iter.Basis.Function(), loc)
		iter.New_root = loc
		iters = append(iters, iter)
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
func interpretSolution(matrix algebra.Matrix) (*base.BasisImpl, float64) {
	b := base.NewBasisImpl(len(matrix)-2)
	last_col := len(matrix[0]) - 1
	for row := 0; row < len(matrix)-1; row++ {
		b.Set(row, matrix[row][last_col])
	}
	return b, math.Abs(matrix[len(matrix)-1][last_col])
}

func updateRoots(roots []float64, orig_func, approx_func Function, loc float64) (replaced float64) {
	// find the first root larger than loc
	i := 0
	for ; i < len(roots); i++ {
		if roots[i] >= loc {
			break;
		}
	}
	max_error := orig_func(loc) - approx_func(loc)
	if i == len(roots) { // loc is after last
		i--
		root_error := orig_func(roots[i]) - approx_func(roots[i])
		if math.Signbit(root_error) != math.Signbit(max_error) {
			i = 0
		}
	} else {
		root_error := orig_func(roots[i]) - approx_func(roots[i])
		if math.Signbit(root_error) != math.Signbit(max_error) {
			if i == 0 { // loc is before first
				// replace the last instead
				i = len(roots)-1
			} else {
				// i is now larger than loc
				// then replace either it or the previous one, depending on sign of errors
				i--
			}
		}
	}
	replaced = roots[i]
	roots[i] = loc
	var tmp sort.Float64Slice = roots
	tmp.Sort()

	return replaced
}
// debug purposes
func printFancy(roots []float64) {
	fmt.Print("[")
	for _, v := range roots {
		fmt.Printf("%.4f ", v)
	}
	fmt.Println("]")
}