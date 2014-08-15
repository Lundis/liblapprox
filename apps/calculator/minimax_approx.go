package main

import (
	"code.google.com/p/liblundis/lmath/exchange"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath"
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"time"
	"strconv"
	"os"
	"fmt"
	"image/color"
	"math"
)

type MinimaxApprox struct {
	id string
	approx *approx.Approx
	iters map[int] []exchange.Iteration
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
	iters := self.iters[degree]
	poly := iters[iter].Poly
	return poly.String()
}

func (self MinimaxApprox) ImageUrl(degree, iter, dimx, dimy int) string {
	filename := self.filename(degree, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	if !ExistsFile(full_path) {
		self.generateImage(degree, iter, dimx, dimy)
	}
	return "img/" + filename
}

func ExistsFile(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func (self MinimaxApprox) generateImage(deg, iter, dimx, dimy int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = fmt.Sprintf("minimax deg %v, iter %v", deg, iter)
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	f := self.approx.Func
	a := self.iters[deg][iter].Poly.Function()

	orig_func := plotter.NewFunction(f)
	orig_func.Color = color.RGBA{R: 255, A: 255}

	approx_func := plotter.NewFunction(a)
	approx_func.Color = color.RGBA{B: 255, A: 255}

	p.Add(orig_func, approx_func)
	p.Legend.Add("f", orig_func)
	p.Legend.Add("p", approx_func)

	p.X.Min = self.approx.Start
	p.X.Max = self.approx.End

	start := self.approx.Start
	end := self.approx.End
	f_min, _ := lmath.Min(f, 100, start, end)
	f_max, _ := lmath.Max(f, 100, start, end)
	a_min, _ := lmath.Min(a, 100, start, end)
	a_max, _ := lmath.Max(a, 100, start, end)
	p.Y.Max = math.Max(f_max, a_max)
	p.Y.Min = math.Min(f_min, a_min)

	file := ImageDir() + string(os.PathSeparator) + self.filename(deg, iter, dimx, dimy)

	if err := p.Save(float64(dimx)/96, float64(dimy)/96, file); err != nil {
		panic(err)
	}
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