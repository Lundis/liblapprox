package approx
import (
	"code.google.com/p/liblundis/lmath/util/cont"
	"code.google.com/p/liblundis/lmath/util/discrete"
    "code.google.com/p/liblundis/lmath/plot"
)

// Common superclass
type ApproxParent struct {
	// Approximated functions of various degrees
	Funcs map[int] cont.Function
	// The max error for each degree. Created by calling PopulateErrors()
	Errors map[int] float64
}

func (self *ApproxParent) init() {
	self.Funcs = make(map[int] cont.Function)
	self.Errors = make(map[int] float64)
}

// Approximation using continuous function
type Approx struct {
	ApproxParent
	// The original function to be approximated
	Func cont.Function
	// Interval endpoints
	Start, End float64
}

func NewApprox(f cont.Function, start, end float64) *Approx {
	approx := new(Approx)
	approx.init()
	approx.Func = f
	approx.Start = start
	approx.End = end
	return approx
}

func (self *Approx) PopulateErrors() {
	self.Errors = make(map[int] float64)
	for degree, f := range self.Funcs {
		self.Errors[degree], _ = cont.FindMaxDiff(self.Func, f, self.Start, self.End)
	}
}

func (self *Approx) SavePlotData(which []int, filename string, points int) {
	plot.WritePlotData(self.Funcs, which, filename, points, self.Start, self.End)
}

// Approximation using discrete points
type DiscreteApprox struct {
	ApproxParent
	// Reference points
	X, Y []float64
}

func NewDiscreteApprox(x, y []float64) *DiscreteApprox {
	approx := new(DiscreteApprox)
	approx.init()
	approx.X = x
	approx.Y = y
	return approx
}

func (self *DiscreteApprox) PopulateErrors() {
	self.Errors = make(map[int] float64)
	for degree, f := range self.Funcs {
		approx_y := discrete.Values(f, self.X)
		self.Errors[degree], _ = discrete.FindMaxDiff(self.Y, approx_y)
	}
}

func (self *DiscreteApprox) SavePlotData(which []int, filename string, points int) {
	plot.WritePlotData(self.Funcs, which, filename, points, self.X[0], self.X[len(self.X)-1])
}