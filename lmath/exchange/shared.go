package exchange

import(
	"code.google.com/p/liblundis/lmath/base"
	"code.google.com/p/liblundis/lmath/algebra"
	"math"
	"fmt"
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

func createMatrix(rows int) algebra.Matrix {
	cols := rows + 1
	return algebra.NewMatrix(rows, cols)
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

// debug purposes
func printFancy(roots []float64) {
	fmt.Print("[")
	for _, v := range roots {
		fmt.Printf("%.4f ", v)
	}
	fmt.Println("]")
}