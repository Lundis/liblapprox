package main

import (
	"code.google.com/p/liblundis/lmath/exchange"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath/plot"
	"code.google.com/p/liblundis/lmath/base/poly"
	"code.google.com/p/liblundis/lmath/util/cont"
	"code.google.com/p/liblundis/lmath/util/discrete"
	"code.google.com/p/gowut/gwu"
	"time"
	"strconv"
	"os"
	"fmt"
)

type MinimaxApproxImpl struct {
	id      string
	iters   map[int] []exchange.Iteration

	info_approx     gwu.Label
	info_replacing  gwu.Label
	info_new_root   gwu.Label
	info_error      gwu.Label
	info_optimality gwu.Label

	ApproxBackendImpl
}

type MinimaxApprox struct {
	MinimaxApproxImpl

	approx *approx.Approx
}

type DiscreteMinimaxApprox struct {
	MinimaxApproxImpl

	approx *approx.DiscreteApprox
}

func NewMinimaxApprox(data *ApproxGUI) *MinimaxApprox {
	ma := new(MinimaxApprox)
	ma.init()
	ma.approx = approx.NewApprox(data.function, data.ival_start, data.ival_end,)
	ma.iters = exchange.Approximate(ma.approx, data.degrees, data.accuracy, poly.PolyFromBasisImpl)
	return ma
}

func NewDiscreteMinimaxApprox(data *ApproxGUI) *DiscreteMinimaxApprox {
	dma := new(DiscreteMinimaxApprox)
	dma.init()
	dma.approx = approx.NewDiscreteApprox(data.discrete_points, discrete.Values(data.function, data.discrete_points))
	dma.iters = exchange.ApproximateDiscrete(dma.approx, data.degrees, poly.PolyFromBasisImpl)
	return dma
}

func (self *MinimaxApproxImpl) init() {
	self.id = "minimax" + strconv.FormatInt(time.Now().UnixNano(), 36)
}

func (self *MinimaxApproxImpl) String(deg, iter int) string {
	iters := self.iters[deg]
	b := iters[iter].Basis
	return b.String()
}

func (self *MinimaxApprox) ImageUrl(deg, iter, dimx, dimy int) string {
	url, exists := imageUrl(self, deg, iter, dimx, dimy)
	if !exists {
		self.generateImage(deg, iter, dimx, dimy)
	}
	return url
}

func (self *DiscreteMinimaxApprox) ImageUrl(deg, iter, dimx, dimy int) string {
	url, exists := imageUrl(self, deg, iter, dimx, dimy)
	if !exists {
		self.generateImage(deg, iter, dimx, dimy)
	}
	return url
}

func (self *MinimaxApprox) generateImage(deg, iter, dimx, dimy int) {
	filepath := ImageDir() + string(os.PathSeparator) + self.Filename(deg, iter, dimx, dimy)
	funcs := []cont.Function{self.approx.Func, self.iters[deg][iter].Basis.Function()}
	labels := []string{"f", "p"}
	title := fmt.Sprintf("minimax deg %v, iter %v", deg, iter)
	plot.SaveSimpleGraph(funcs, labels, self.approx.Start, self.approx.End, title, filepath, dimx, dimy)
}

func (self *DiscreteMinimaxApprox) generateImage(deg, iter, dimx, dimy int) {
	filepath := ImageDir() + string(os.PathSeparator) + self.Filename(deg, iter, dimx, dimy)
	points_x := [][]float64{self.approx.X, self.approx.X}
	approx_y := discrete.Values(self.iters[deg][iter].Basis.Function(), self.approx.X)
	points_y := [][]float64{self.approx.Y, approx_y}
	labels := []string{"f", "p"}
	title := fmt.Sprintf("minimax deg %v, iter %v", deg, iter)
	plot.SaveSimpleGraphXY(points_x, points_y, labels, title, filepath, dimx, dimy)
	
}

func (self *MinimaxApprox) ErrorGraphUrl(deg, iter, dimx, dimy int) string {
	filename := "err_" + self.Filename(deg, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	if !existsFile(full_path) {
		self.generateErrorGraph(deg, iter, dimx, dimy)
	}
	return "img/" + filename
}

func (self *DiscreteMinimaxApprox) ErrorGraphUrl(deg, iter, dimx, dimy int) string {
	filename := "err_" + self.Filename(deg, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	if !existsFile(full_path) {
		self.generateErrorGraph(deg, iter, dimx, dimy)
	}
	return "img/" + filename
}

func (self *MinimaxApprox) generateErrorGraph(deg, iter, dimx, dimy int) {
	filepath := ImageDir() + string(os.PathSeparator) + "err_" + self.Filename(deg, iter, dimx, dimy)
	f := func(x float64) float64 {
		return self.approx.Func(x) - self.iters[deg][iter].Basis.Function()(x)
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
	funcs := []cont.Function{z, f, hp, hm}
	labels := []string{"0", "e = f - p", "h", "-h"}
	title := fmt.Sprintf("error deg %v, iter %v", deg, iter)
	plot.SaveSimpleGraph(funcs, labels, self.approx.Start, self.approx.End, title, filepath, dimx, dimy)
}

func (self *DiscreteMinimaxApprox) generateErrorGraph(deg, iter, dimx, dimy int) {
	filepath := ImageDir() + string(os.PathSeparator) + "err_" + self.Filename(deg, iter, dimx, dimy)

	x := self.approx.X

	// elements of z init by default to 0
	z := make([]float64, len(x))

	hp := make([]float64, len(x))
	hm := make([]float64, len(x))
	h := self.iters[deg][iter].Leveled_error
	for i := range hp {
		hp[i] = h
		hm[i] = -h
	}

	approx_y := discrete.Values(self.iters[deg][iter].Basis.Function(), x)
	error := discrete.Minus(self.approx.Y, approx_y)

	points_x := [][]float64{x, x, x, x}
	points_y := [][]float64{z, error, hp, hm}


	labels := []string{"0", "e = f - p", "h", "-h"}
	title := fmt.Sprintf("error deg %v, iter %v", deg, iter)
	plot.SaveSimpleGraphXY(points_x, points_y, labels, title, filepath, dimx, dimy)
}

func (self *MinimaxApproxImpl) Error(deg, iter int) float64 {
	return self.iters[deg][iter].Max_error
}

func (self *MinimaxApproxImpl) Optimality(deg, iter int) float64 {
	return self.iters[deg][iter].ErrorDiff()
}

func (self *MinimaxApproxImpl) Iters(deg int) int {
	return len(self.iters[deg])
}

func (self *MinimaxApproxImpl) BuildInfoTable() gwu.Comp {
	t := gwu.NewTable()
	t.SetCellPadding(4)
	t.Style().SetBorder2(1, gwu.BRD_STYLE_SOLID, gwu.CLR_BLACK)
	t.EnsureSize(4, 2)

	t.Add(gwu.NewLabel("Approximation"), 0, 0)
	self.info_approx = gwu.NewLabel("")
	t.Add(self.info_approx, 0, 1)

	t.Add(gwu.NewLabel("Replacing"), 1, 0)
	self.info_replacing = gwu.NewLabel("")
	t.Add(self.info_replacing, 1, 1)

	t.Add(gwu.NewLabel("New root"), 2, 0)
	self.info_new_root = gwu.NewLabel("")
	t.Add(self.info_new_root, 2, 1)

	t.Add(gwu.NewLabel("Max Error"), 3, 0)
	self.info_error = gwu.NewLabel("")
	t.Add(self.info_error, 3, 1)

	t.Add(gwu.NewLabel("Optimality"), 4, 0)
	self.info_optimality = gwu.NewLabel("")
	t.Add(self.info_optimality, 4, 1)
	return t
}

func (self *MinimaxApproxImpl) UpdateInfoTable(deg, iter int) {
	self.info_approx.SetText(self.String(deg, iter))
	self.info_replacing.SetText(fmt.Sprintf("%v", self.iters[deg][iter].Replacing))
	self.info_new_root.SetText(fmt.Sprintf("%v", self.iters[deg][iter].New_root))
	self.info_error.SetText(fmt.Sprintf("%v", self.Error(deg, iter)))
	self.info_optimality.SetText(fmt.Sprintf("%v", self.Optimality(deg, iter)))
}