package main

import (
	"code.google.com/p/liblundis/lmath/exchange"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath/plot"
	"code.google.com/p/liblundis/lmath"
	//"code.google.com/p/plotinum/plot"
	//"code.google.com/p/plotinum/plotter"
	"time"
	"strconv"
	"os"
	"fmt"
)

type MinimaxApprox struct {
	id string
	approx *approx.Approx
	iters map[int] []exchange.Iteration
}

func NewMinimaxApprox(data *ApproxGUI) *MinimaxApprox {
	return newMinimaxApprox(data.degrees, data.function, data.ival_start, data.ival_end, data.accuracy)
}

func newMinimaxApprox(degs []int, f lmath.Function, start, end, accuracy float64) *MinimaxApprox {
	mma := new(MinimaxApprox)
	mma.id = "minimax" + strconv.FormatInt(time.Now().UnixNano(), 36)
	mma.approx = approx.NewApprox(f, start, end)
	mma.iters = exchange.Approximate(mma.approx, degs, accuracy)
	return mma
}

func (self MinimaxApprox) String(deg, iter int) string {
	iters := self.iters[deg]
	poly := iters[iter].Poly
	return poly.String()
}

func (self MinimaxApprox) ImageUrl(deg, iter, dimx, dimy int) string {
	filename := self.filename(deg, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	if !ExistsFile(full_path) {
		self.generateImage(deg, iter, dimx, dimy)
	}
	return "img/" + filename
}

func ExistsFile(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func (self MinimaxApprox) generateImage(deg, iter, dimx, dimy int) {
	filepath := ImageDir() + string(os.PathSeparator) + self.filename(deg, iter, dimx, dimy)
	funcs := []lmath.Function{self.approx.Func, self.iters[deg][iter].Poly.Function()}
	labels := []string{"f", "p"}
	title := fmt.Sprintf("minimax deg %v, iter %v", deg, iter)
	plot.SaveSimpleGraph(funcs, labels, self.approx.Start, self.approx.End, title, filepath, dimx, dimy)
}

func (self MinimaxApprox) ErrorGraphUrl(deg, iter, dimx, dimy int) string {
	filename := "err_" + self.filename(deg, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	if !ExistsFile(full_path) {
		self.generateErrorGraph(deg, iter, dimx, dimy)
	}
	return "img/" + filename
}

func (self MinimaxApprox) generateErrorGraph(deg, iter, dimx, dimy int) {
	filepath := ImageDir() + string(os.PathSeparator) + "err_" + self.filename(deg, iter, dimx, dimy)
	f := func(x float64) float64 {
		return self.approx.Func(x) - self.iters[deg][iter].Poly.Function()(x)
	}
	z := func(x float64) float64 {
		return 0
	}
	h := self.iters[deg][iter].Leveled_error
	hp := func(x float64) float64 {
		return h
	}
	hm := func(x float64) float64 {
		return -h
	}
	funcs := []lmath.Function{z, f, hp, hm}
	labels := []string{"0", "e = f - p", "h", "-h"}
	title := fmt.Sprintf("error deg %v, iter %v", deg, iter)
	plot.SaveSimpleGraph(funcs, labels, self.approx.Start, self.approx.End, title, filepath, dimx, dimy)
}

func (self MinimaxApprox) filename(deg, iter, dimx, dimy int) string {
	return fmt.Sprintf("%v_%v_%v_%v_%v.png", self.id, deg, iter, dimx, dimy)
}

func (self MinimaxApprox) Error(deg, iter int) float64 {
	return self.iters[deg][iter].Max_error
}

func (self MinimaxApprox) Optimality(deg, iter int) float64 {
	return self.iters[deg][iter].ErrorDiff()
}

func (self MinimaxApprox) Iters(deg int) int {
	return len(self.iters[deg])
}