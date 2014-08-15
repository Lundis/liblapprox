package main

import (
	"code.google.com/p/liblundis/lmath"
	"code.google.com/p/liblundis/lmath/algebra"
	"code.google.com/p/gowut/gwu"
	"strings"
	"strconv"
	"errors"
	"math"
	"fmt"
)

type ApproxGUI struct {
	function_box         gwu.ListBox
	function             lmath.Func1to1
	interval_box         gwu.TextBox
	ival_start, ival_end float64

	degrees_box   gwu.TextBox
	degrees     []int
	current_deg   int
	current_iter  int

	approx_box   gwu.ListBox
	accuracy_box gwu.TextBox
	accuracy     float64

	result_container  gwu.Panel
	degree_buttons    gwu.Panel
	iter_buttons      gwu.Panel

	dimx_box  gwu.TextBox
	dimy_box  gwu.TextBox
	dimx      int
	dimy      int
	image     gwu.Image

	err               error
	backend           ApproxBackend

	message         gwu.Label
	info_approx     gwu.Label
	info_degree     gwu.Label
	info_iter       gwu.Label
	info_error      gwu.Label
	info_optimality gwu.Label
}

func (self *ApproxGUI) showError(m string) {
	self.message.SetText(m)
	self.message.Style().SetColor(gwu.CLR_RED)
}

func (self *ApproxGUI) showMessage(m string) {
	self.message.SetText(m)
	self.message.Style().SetColor(gwu.CLR_BLACK)
}

func (self *ApproxGUI) clearMessage() {
	self.message.SetText("")
	self.message.Style().SetColor(gwu.CLR_BLACK)
}


func (self *ApproxGUI) updateResultBrowser() {
	self.clearMessage()
	if self.err != nil {
		self.showError(self.err.Error())
	} else if self.backend == nil {
		self.showMessage("Please select a function and approximation and click on the button!")
	}
	self.updateDegreeSelector()
	self.updateIterationSelector()
	self.updateGraphViewer()
	self.updateInfoTable()
}

func (self *ApproxGUI) updateDegreeSelector() {
	self.degree_buttons.Clear()
	if self.backend != nil && self.err == nil {
		for i := 0; i < len(self.degrees); i++ {
			deg := self.degrees[i]
			b := gwu.NewButton(fmt.Sprintf("%v", deg))
			if self.current_deg == deg {
				b.SetEnabled(false)
			} else {
				b.AddEHandlerFunc(func(ev gwu.Event) {
					self.current_deg = deg
					self.current_iter = 0
					self.updateResultBrowser()
					ev.MarkDirty(self.result_container)
				}, gwu.ETYPE_CLICK)
			}
			self.degree_buttons.Add(b)
		}
	}
}

func (self *ApproxGUI) updateIterationSelector() {
	self.iter_buttons.Clear()
	if self.backend != nil && self.err == nil {
		for i := 0; i < self.backend.Iters(self.current_deg); i++ {
			i_local := i // this is required for the handler function to get the right i
			b := gwu.NewButton(fmt.Sprintf("%v", i))
			if self.current_iter == i {
				b.SetEnabled(false)
			} else {
				b.AddEHandlerFunc(func(ev gwu.Event) {
					self.current_iter = i_local
					self.updateResultBrowser()
					ev.MarkDirty(self.result_container)
				}, gwu.ETYPE_CLICK)
			}
			self.iter_buttons.Add(b)
		}
	}
}

func (self *ApproxGUI) updateGraphViewer() {
	self.image.SetUrl("")
	if self.backend != nil && self.err == nil {
		self.image.SetUrl(self.backend.ImageUrl(self.current_deg, self.current_iter, self.dimx, self.dimy))
	}
}

func (self *ApproxGUI) updateInfoTable() {
	if self.backend != nil && self.err == nil {
		deg := self.current_deg
		iter := self.current_iter
		self.info_approx.SetText(self.backend.String(deg, iter))
		self.info_degree.SetText(fmt.Sprintf("%v", deg))
		self.info_iter.SetText(fmt.Sprintf("%v", iter))
		self.info_error.SetText(fmt.Sprintf("%v", self.backend.Error(deg, iter)))
		self.info_optimality.SetText(fmt.Sprintf("%v", self.backend.Optimality(deg, iter)))
	}
	
}

func (self *ApproxGUI) interpretImageDim() error {
	strx := strings.Replace(self.dimx_box.Text(), " ", "", -1)
	stry := strings.Replace(self.dimy_box.Text(), " ", "", -1)
	x, err := strconv.ParseInt(strx, 10, 32)
	if err != nil {
		self.err = err
		return err
	}
	y, err := strconv.ParseInt(stry, 10, 32)
	if err != nil {
		self.err = err
		return err
	}
	if x < 100 || y < 100 || x > 2000 || y > 2000 {
		self.err = errors.New("Graph size must be between 100x100px and 2000x2000px")
		return self.err
	}
	self.dimx = int(x)
	self.dimy = int(y)
	return nil
}

func (self *ApproxGUI) interpretSettings() error {
	self.err = self.interpretDegrees()
	if self.err != nil {
		return self.err
	}
	self.err = self.interpretIntervals()
	if self.err != nil {
		return self.err
	}
	self.err = self.interpretAccuracy()
	if self.err != nil {
		return self.err
	}
	self.interpretFunction()
	return nil
}

func (self *ApproxGUI) interpretDegrees() error {
	text := strings.Replace(self.degrees_box.Text(), " ", "", -1)
	split := strings.Split(text, ",")
	self.degrees = make([]int, 0, 10)
	for _, str := range split {
		err := self.interpretDegree(str)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *ApproxGUI) interpretDegree(num string) error {
	if hyph := strings.Index(num, "-"); hyph != -1 {
		first, err1 := strconv.ParseInt(num[:hyph], 10, 16)
		if err1 != nil {
			return err1
		}
		second, err2 := strconv.ParseInt(num[hyph+1:], 10, 16)
		if err2 != nil {
			return err2
		}
		if first >= second || first < 0 {
			return errors.New(fmt.Sprintf("first value (%v) in range must be >= 0 and < second(%v)", first, second))
		} else {
			for i := int(first); i <= int(second); i++ {
				self.degrees = append(self.degrees, i)
			}
		}
	} else {
		int_, err := strconv.ParseInt(num, 10, 16)
		if err != nil {
			return err
		}
		self.degrees = append(self.degrees, int(int_))
	}
	
	return nil
}

func (self *ApproxGUI) interpretIntervals() error {
	str := strings.Replace(self.interval_box.Text(), " ", "", -1)
	split := strings.Split(str, ",")
	if len(split) != 2 {
		return errors.New("interval endpoints must be separated by a comma")
	}
	first, err := algebra.EvalCommonf64(split[0])
	if err != nil {
		return err
	}
	second, err := algebra.EvalCommonf64(split[1])
	if err != nil {
		return err
	}
	self.ival_start = first
	self.ival_end = second
	return nil
}

// checks the function dropdown menu for its selection and 
// creates the actual function used for the approximation
func (self *ApproxGUI) interpretFunction() {
	switch self.function_box.SelectedValue() {
	case "sin x":
		self.function = func(x float64) float64 {
			return math.Sin(x)
		}
	case "cos x":
		self.function = func(x float64) float64 {
			return math.Cos(x)
		}
	case "e^x":
		self.function = func(x float64) float64 {
			return math.Pow(math.E, x)
		}
	case "1/sqrt(2*pi) * e^(-x^2 / 2":
		self.function = func(x float64) float64 {
			return 1.0 / math.Sqrt(2*math.Pi) * math.Pow(math.E, -x*x/2)
		}
	default:
		panic(fmt.Sprintf("unknown function selected: %v", self.function_box.SelectedValue()))
	}
}

func (self *ApproxGUI) interpretAccuracy() error {
	a, err := strconv.ParseFloat(self.accuracy_box.Text(), 64)
	if err != nil {
		self.err = errors.New("Error parsing accuracy")
		return self.err
	} else {
		self.accuracy = a
		return nil
	}
}

func (self *ApproxGUI) approximateMinimax() {
	self.backend = NewMinimaxApprox(self)
	self.current_deg = self.degrees[0]
	self.current_iter = 0
}