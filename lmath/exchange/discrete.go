package exchange

import (
	"code.google.com/p/liblundis/lmath/util/cont"
	. "code.google.com/p/liblundis/lmath/util/discrete"
	"code.google.com/p/liblundis/lmath/base"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath/algebra"
	"math"
	"fmt"
	"sort"
)

// One big difference between this implementation and the continous one is that in this one the roots slice actually contains indices and not values directly


// Approximates for the specified degrees, saving the result in approx.
// Returns details about the iterations.
func ApproximateDiscrete(approx *approx.DiscreteApprox, degrees []int, basis_func base.BasisImplConverter) map[int] []Iteration {
	iters := make(map[int] []Iteration)
	for _, deg := range degrees {
		iters[deg] = ApproximateDegreeDiscrete(approx, deg, basis_func)
	}
	return iters
}

// Approximates one degree, saving the result in approx
// Returns information about the iterations
func ApproximateDegreeDiscrete(approx *approx.DiscreteApprox, degree int, basis_func base.BasisImplConverter) []Iteration {
	iters := make([]Iteration, 0, 10)
	roots := make([]int, degree + 2)
	for i := range roots {
		roots[i] = i
	}

	matrix := createMatrix(len(roots))
	old_leveled_error := float64(0)
	for {
		updateMatrixDiscrete(matrix, approx, roots)
		matrix.Solve()
		iter := iterationDiscrete(matrix, approx, roots, basis_func)
		iters = append(iters, iter)
		if old_leveled_error >= iter.Leveled_error {
			fmt.Printf("exchange failed to converge for deg %v!\n", degree)
			// if the new leveled error is smaller than the old the algorithm is broken and should stop
			break
		} else {
			old_leveled_error = iter.Leveled_error
		}
	}
	approx.Funcs[degree] = iters[len(iters)-1].Basis.Function()
	approx.Errors[degree] = iters[len(iters)-1].Max_error
	return iters
}

func updateMatrixDiscrete(m algebra.Matrix, approx *approx.DiscreteApprox, roots []int) {
	degree := len(roots)-2
	for row, equation := range m {
		x_index := roots[row]
		fmt.Printf("", x_index, approx.X[x_index])
		for col := 0; col <= degree; col++ {
			equation[col] = math.Pow(approx.X[x_index], float64(col))
		}
		// error with alternating signs
		equation[degree+1] = math.Pow(-1, float64(row % 2))
		// actual function value
		equation[degree+2] = approx.Y[x_index]
	}
}

func iterationDiscrete(matrix algebra.Matrix, approx *approx.DiscreteApprox, roots []int, basis_func base.BasisImplConverter) Iteration {
	iter := Iteration{}
	var b *base.BasisImpl
	b, iter.Leveled_error = interpretSolution(matrix)
	iter.Basis = basis_func(b)

	var new_root_index int
	iter.Max_error, new_root_index = FindMaxDiff(approx.Y, Values(iter.Basis.Function(), approx.X))
	iter.Replacing = updateRootsDiscrete(roots, approx, iter.Basis.Function(), new_root_index)
	iter.New_root = approx.X[new_root_index]
	return iter
}

// index refers to the discrete index where the max error was found
func updateRootsDiscrete(roots []int, approx *approx.DiscreteApprox, approx_func cont.Function, new_root_index int) (replaced float64) {
	
	new_root := approx.X[new_root_index]
	// find the first root larger than the new one
	old_root_index := 0 // index in roots
	for ; old_root_index < len(roots); old_root_index++ {
		if approx.X[roots[old_root_index]] >= new_root {
			break
		}
	}

	max_error := approx.Y[new_root_index] - approx_func(approx.X[new_root_index])

	if old_root_index == len(roots) { // new_root is after last
		old_root_index--
		old_root_x := approx.X[roots[old_root_index]]
		root_error := approx.Y[new_root_index] - approx_func(old_root_x)
		if math.Signbit(root_error) != math.Signbit(max_error) {
			old_root_index = 0
		}
	} else {
		old_root_x := approx.X[roots[old_root_index]]
		root_error := approx.Y[new_root_index] - approx_func(old_root_x)
		if math.Signbit(root_error) != math.Signbit(max_error) {
			if old_root_index == 0 { // new_root is before first
				// replace the last instead
				old_root_index = len(roots)-1
			} else {
				// i is now larger than new_root
				// then replace either it or the previous one, depending on sign of errors
				old_root_index--
			}
		}
	}
	replaced = approx.Y[roots[old_root_index]]
	roots[old_root_index] = new_root_index
	var tmp sort.IntSlice = roots
	tmp.Sort()

	return replaced
}
