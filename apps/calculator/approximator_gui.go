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

	result_container  gwu.Panel
	iteration_browser gwu.Panel


	dimx_box  gwu.TextBox
	dimy_box  gwu.TextBox
	dimx      int
	dimy      int
	image     gwu.Image

	err               error
	backend           ApproxBackend
	accuracy          float64

	info_degree     gwu.Label
	info_iter       gwu.Label
	info_error      gwu.Label
	info_optimality gwu.Label
}

func (self *ApproxGUI) BuildGUI() gwu.Comp {
	p := gwu.NewVerticalPanel()
	p.Add(self.buildConfigSection())
	self.result_container = gwu.NewPanel()
	self.updateResultBrowser()
	p.Add(self.result_container)
	return p
}

func (self *ApproxGUI) buildConfigSection() gwu.Comp {
	p := gwu.NewVerticalPanel()
	function_selector := self.buildFunctionSelector()
	p.Add(function_selector)
	approx_settings := self.buildApproximationSelector()
	p.Add(approx_settings)
	return p
}

func (self *ApproxGUI) buildFunctionSelector() gwu.Comp {
	function_selector := gwu.NewHorizontalPanel()
	function_selector.Add(gwu.NewLabel("Function:"))
	lb := gwu.NewListBox([]string{"sin x", "cos x", "e^x", "1/sqrt(2*pi) * e^(-x^2 / 2"})
	lb.SetSelected(0, true)
	self.function_box = lb
	function_selector.Add(self.function_box)
	function_selector.Add(gwu.NewLabel("Intervals (separated by comma):"))
	self.interval_box = gwu.NewTextBox("0,1")
	function_selector.Add(self.interval_box)
	return function_selector
}

func (self *ApproxGUI) buildApproximationSelector() gwu.Comp {
	approx_settings := gwu.NewHorizontalPanel()
	approx_type := gwu.NewListBox([]string{"minimax"})
	approx_type.SetSelected(0, true)
	self.approx_box = approx_type
	approx_settings.Add(approx_type)
	degbox := gwu.NewTextBox("1,2,3")
	self.degrees_box = degbox
	self.degrees_box.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	approx_settings.Add(self.degrees_box)

	run_button := gwu.NewButton("Approximate!")
	run_button.AddEHandlerFunc(func(ev gwu.Event) {
		err := self.interpretSettings()
		if err == nil {
			switch self.approx_box.SelectedValue() {
			case "minimax":
				self.approximateMinimax()
			default:
				panic(fmt.Sprintf("Unknown approximation type selected: %v", self.approx_box.SelectedValue()))
			}
		}
		self.updateResultBrowser()
		ev.MarkDirty(self.result_container)
	}, gwu.ETYPE_CLICK)
	approx_settings.Add(run_button)
	return approx_settings
}

func (self *ApproxGUI) updateResultBrowser() {
	self.result_container.Clear()
	if self.err != nil {
		err := gwu.NewLabel(self.err.Error())
		err.Style().SetColor(gwu.CLR_RED)
		self.result_container.Add(err)
	} 
	if self.backend == nil {
		l := gwu.NewLabel("Please select a function and approximation and click on the button!")
		self.result_container.Add(l)
	} else {
		p := gwu.NewVerticalPanel()
		p.Add(self.buildDegreeSelector())
		p.Add(self.buildIterationSelector())
		self.result_container.Add(p)
	}
}

func (self *ApproxGUI) buildDegreeSelector() gwu.Panel {
	p := gwu.NewNaturalPanel()
	l := gwu.NewLabel("Degree")
	p.Add(l)
	for i := 0; i < len(self.degrees); i++ {
		b := gwu.NewButton(fmt.Sprintf("%v", self.degrees[i]))
		b.AddEHandlerFunc(func(ev gwu.Event) {

			ev.MarkDirty(self.iteration_browser)
		}, gwu.ETYPE_CLICK)
		p.Add(b)
	}
	return p
}

func (self *ApproxGUI) buildIterationBrowser() gwu.Panel {
	p := gwu.NewVerticalPanel()
	p.Add(self.buildIterationSelector())
	hp := gwu.NewHorizontalPanel()
	hp.Add(self.buildGraphViewer())
	hp.Add(self.buildInfoViewer())
	p.Add(hp)
	return p
}

func (self *ApproxGUI) buildIterationSelector() gwu.Panel {
	p := gwu.NewNaturalPanel()
	l := gwu.NewLabel("Iteration")
	p.Add(l)
	for i := 0; i < self.backend.Iters(self.current_deg); i++ {
		b := gwu.NewButton(fmt.Sprintf("%v", self.degrees[i]))
		b.AddEHandlerFunc(func(ev gwu.Event) {
			self.current_iter = i
			self.refreshImage(ev)
		}, gwu.ETYPE_CLICK)
		p.Add(b)
	}
	self.iteration_browser = p
	return p
}

func (self *ApproxGUI) buildGraphViewer() gwu.Panel {
	hp := gwu.NewHorizontalPanel()
	vp := gwu.NewVerticalPanel()
	l := gwu.NewLabel("Graph dimensions: ")
	vp.Add(l)
	self.dimx_box = gwu.NewTextBox("600")
	self.dimx = 600
	vp.Add(self.dimx_box)
	self.dimy_box = gwu.NewTextBox("400")
	self.dimy = 400
	vp.Add(self.dimy_box)
	b := gwu.NewButton("set")
	b.AddEHandlerFunc(func(ev gwu.Event) {
		if err := self.interpretImageDim(); err != nil {
			self.refreshImage(ev)
		}
	}, gwu.ETYPE_CLICK)
	vp.Add(b)
	hp.Add(vp)
	self.image = gwu.NewImage("approximation graph", "")
	hp.Add(self.image)
	return hp
}

// shows the currently selected approx/iter/size
func (self *ApproxGUI) refreshImage(ev gwu.Event) {
	self.image.SetUrl(self.backend.ImageUrl(self.current_deg, self.current_iter, self.dimx, self.dimy))
	ev.MarkDirty(self.image)
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

func (self *ApproxGUI) buildInfoViewer() gwu.Comp {
	t := gwu.NewTable()
	t.EnsureSize(4, 2)
	t.Add(gwu.NewLabel("Degree"), 0, 0)
	self.info_degree = gwu.NewLabel(fmt.Sprintf("%v", self.current_deg))
	t.Add(self.info_degree, 0, 1)

	t.Add(gwu.NewLabel("Iteration"), 1, 0)
	self.info_iter = gwu.NewLabel(fmt.Sprintf("%v", self.current_iter))
	t.Add(self.info_iter, 1, 1)

	t.Add(gwu.NewLabel("Max Error"), 2, 0)
	self.info_error = gwu.NewLabel(fmt.Sprintf("%v", -1))
	t.Add(self.info_error, 2, 1)

	t.Add(gwu.NewLabel("Optimality"), 3, 0)
	self.info_optimality = gwu.NewLabel(fmt.Sprintf("%v", -1))
	t.Add(self.info_optimality, 3, 1)
	return t
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

func (self *ApproxGUI) approximateMinimax() {
	self.backend = NewMinimaxApprox(self)
}