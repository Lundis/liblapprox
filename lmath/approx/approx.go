package approx
import (
	. "code.google.com/p/liblundis/lmath"
    "code.google.com/p/liblundis/lmath/plot"
)

// Approx contains
type Approx struct {
	// The original function to be approximated
	Func Function
	// Approximated functions of various degrees
	Funcs map[int] Function
	// The max error for each degree. Created by calling PopulateErrors()
	Errors map[int] float64
	// Interval endpoints
	Start, End float64
}

func NewApprox(f Function, start, end float64) *Approx {
	approx := new(Approx)
	approx.Funcs = make(map[int] Function)
	approx.Func = f
	approx.Start = start
	approx.End = end
	return approx
}

func (self *Approx) PopulateErrors() {
	self.Errors = make(map[int] float64)
	for degree, f := range self.Funcs {
		self.Errors[degree], _ = FindMaxDiff(self.Func, f, self.Start, self.End)
	}
}

func (self *Approx) SavePlotData(which []int, filename string, points int) {
	plot.WritePlotData(self.Funcs, which, filename, points, self.Start, self.End)
}