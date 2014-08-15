package main

import (
	"code.google.com/p/liblundis/lmath/exchange"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath"
	"time"
	"strconv"
	"os"
	"fmt"
)

type MinimaxApprox struct {
	id string
	approx *approx.Approx
	iters map[int][]exchange.Iteration
}

func NewMinimaxApprox(data *ApproxGUI) *MinimaxApprox {
	return newMinimaxApprox(data.degrees, data.function, data.ival_start, data.ival_end, data.accuracy)
}

func newMinimaxApprox(degrees []int, f lmath.Func1to1, start, end, accuracy float64) *MinimaxApprox {
	mma := new(MinimaxApprox)
	mma.id = "minimax" + strconv.FormatInt(time.Now().UnixNano(), 36)
	mma.approx = approx.NewApprox(f, start, end)
	mma.iters = exchange.Approximate(mma.approx, degrees, accuracy)
	return mma
}

func (self MinimaxApprox) String(degree, iter int) string {
	return self.iters[degree][iter].Poly.String()
}

func (self MinimaxApprox) ImageUrl(degree, iter, dimx, dimy int) string {
	filename := self.filename(degree, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	if !ExistsFile(full_path) {
		self.generateImage(degree, iter, dimx, dimy)
	}
	return filename
}

func ExistsFile(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func (self MinimaxApprox) generateImage(degree, iter, dimx, dimy int) {

}

func (self MinimaxApprox) filename(degree, iter, dimx, dimy int) string {
	return fmt.Sprintf("%v_%v_%v_%v_%v.png", self.id, degree, iter, dimx, dimy)
}

func (self MinimaxApprox) Error(degree, iter int) float64 {
	return self.iters[degree][iter].Max_error
}

func (self MinimaxApprox) Optimality(degree, iter int) float64 {
	return self.iters[degree][iter].ErrorDiff()
}

func (self MinimaxApprox) Iters(degree int) int {
	return len(self.iters[degree])
}