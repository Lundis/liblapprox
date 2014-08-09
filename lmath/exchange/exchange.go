package exchange

import (
	"code.google.com/p/liblundis/lmath"
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
	return fmt.Sprintf("%v || error: %v", self.Poly, self.Max_error - self.Leveled_error)
}

// Approximates for the specified degrees, saving the result in approx.
// Returns details about the iterations.
func Approximate(approx *approx.Approx, degrees []int, accuracy float64) map[int] []Iteration {
	iters := make(map[int] []Iteration, len(degrees))
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
	diff := accuracy*2
	for diff > accuracy {
		matrix.update(approx, roots)
		matrix.solve()
		iter := Iteration{}
		iter.Poly, iter.Leveled_error = matrix.interpretSolution()
		var loc float64
		iter.Max_error, loc = lmath.FindMaxDiff(approx.Func, iter.Poly.Function(), approx.Start, approx.End)
		diff = iter.Max_error - iter.Leveled_error
		iters = append(iters, iter)
		updateRoots(roots, iter.Poly.Function(), loc, iter.Max_error)
	}
	

	return iters
}

type matrix [][]float64

func createMatrix(approx *approx.Approx, roots []float64) matrix {
	degree := len(roots)-2
	matrix := make([][]float64, len(roots))
	for row := range matrix {
		// degree n has n+1 columns
		// + one col for error, one for f(x)
		matrix[row] = make([]float64, degree + 3)
	}
	return matrix
}

func (self matrix) update(approx *approx.Approx, roots []float64) {
	degree := len(roots)-2
	for row, equation := range self {
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

func (self matrix) solve() {
	for col := 0; col < len(self[0]) - 1; col++ {
		self.ensureDiagonal1(col)
		self.reduceAllBut(col)
	}
}

// Make sure that the leading number of row col is 1, by dividing it with itself.
// If it is zero, swap it with one further down, then do it again
func (self matrix) ensureDiagonal1(col int) {
	if self[col][col] != 0 {
		div := self[col][col]
		for i := col; i < len(self[col]); i++ {
			self[col][i] /= div
		}
	} else {
		for i := col; i < len(self); i++ {
			if self[i][col] != 0 {
				self.swapRows(i, col)
				self.ensureDiagonal1(col)
			}
		}
	}
}

func (self matrix) swapRows(r1, r2 int) {
	tmp := self[r1]
	self[r1] = self[r2]
	self[r2] = tmp
}

// subtracts rows with each other so that everything in column col but matrix[col][col] is 0
func (self matrix) reduceAllBut(col int) {
	for row := 0; row < len(self); row++ {
		if row == col {
			continue
		}
		// calculate the multiplier
		mult := self[row][col]
		self[row][col] = 0
		for i := col + 1; i < len(self[row]); i++ {
			self[row][i] -= mult*self[col][i]
		}
	}
}

// interprets the solution and returns the polynomial and the leveled error
func (self matrix) interpretSolution() (lmath.Polynomial, float64) {
	poly := lmath.NewPolynomial(len(self)-2)
	last_col := len(self[0]) - 1
	for row := 0; row < len(self)-1; row++ {
		poly[row] = self[row][last_col]
	}
	return poly, math.Abs(self[len(self)-1][last_col])
}

func updateRoots(roots []float64, approx_func lmath.Func1to1, loc, max_error float64) {
	if len(roots) == 1 {
		roots[0] = loc
		return
	}
	// find the first root larger than loc
	i := 0
	for ; i < len(roots)-1; i++ {
		if roots[i] > loc {
			break;
		}
	}
	// then replace either it or the previous/next one, depending on sign of errors
	if approx_func(roots[i]) < 0 {
		if approx_func(loc) < 0 {
			roots[i] = loc
		} else {
			roots[i-1] = loc
		}
	} else {
		if approx_func(loc) > 0 {
			if i > 0 {
				roots[i-1] = loc
			} else {
				roots[i+1] = loc
			}
			
		} else {
			roots[i] = loc
		}
	}

}