package main

import (
	"code.google.com/p/gowut/gwu"
	"fmt"
)


func (self *ApproxGUI) BuildGUI() gwu.Comp {
	p := gwu.NewVerticalPanel()
	p.Add(self.buildConfigSection())
	self.result_container = gwu.NewPanel()
	self.buildResultBrowser()
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
			fmt.Println("Approximating...")
			switch self.approx_box.SelectedValue() {
			case "minimax":
				self.approximateMinimax()
			default:
				panic(fmt.Sprintf("Unknown approximation type selected: %v", self.approx_box.SelectedValue()))
			}
		}
		fmt.Println("done!")
		self.updateResultBrowser()
		ev.MarkDirty(self.result_container)
	}, gwu.ETYPE_CLICK)
	approx_settings.Add(gwu.NewLabel("Accuracy: "))
	self.accuracy_box = gwu.NewTextBox("1e-7")
	self.accuracy = 1e-7
	approx_settings.Add(self.accuracy_box)
	approx_settings.Add(run_button)
	return approx_settings
}

func (self *ApproxGUI) buildResultBrowser() {
	self.message = gwu.NewLabel("Please select a function and approximation and click on the button!")
	self.result_container.Add(self.message)

	p := gwu.NewVerticalPanel()
	p.Add(self.buildDegreeSelector())
	p.Add(self.buildIterationSelector())
	p.Add(self.buildGraphSizeSelector())
	p.Add(self.buildGraphViewer())
	p.Add(self.buildInfoViewer())
	p.Add(self.buildErrorTable())
	self.result_container.Add(p)
}

func (self *ApproxGUI) buildDegreeSelector() gwu.Panel {
	p := gwu.NewVerticalPanel()
	l := gwu.NewLabel("Degree")
	p.Add(l)
	self.degree_buttons = gwu.NewNaturalPanel()
	p.Add(self.degree_buttons)
	return p
}

func (self *ApproxGUI) buildIterationSelector() gwu.Panel {
	p := gwu.NewHorizontalPanel()
	l := gwu.NewLabel("Iteration")
	p.Add(l)
	self.iter_buttons = gwu.NewNaturalPanel()
	p.Add(self.iter_buttons)
	return p
}

func (self *ApproxGUI) buildGraphSizeSelector() gwu.Panel {
	hp := gwu.NewHorizontalPanel()
	l := gwu.NewLabel("Graph dimensions: ")
	hp.Add(l)
	self.dimx_box = gwu.NewTextBox("600")
	self.dimx = 600
	hp.Add(self.dimx_box)
	self.dimy_box = gwu.NewTextBox("400")
	self.dimy = 400
	hp.Add(self.dimy_box)
	b := gwu.NewButton("set")
	b.AddEHandlerFunc(func(ev gwu.Event) {
		self.interpretImageDim()
		self.updateResultBrowser()
		ev.MarkDirty(self.result_container)
	}, gwu.ETYPE_CLICK)
	hp.Add(b)
	return hp
}

func (self *ApproxGUI) buildGraphViewer() gwu.Panel {
	p := gwu.NewHorizontalPanel()
	self.image = gwu.NewImage("approximation graph", "")
	p.Add(self.image)
	self.error_image = gwu.NewImage("error graph", "")
	p.Add(self.error_image)
	return p
}

func (self *ApproxGUI) buildInfoViewer() gwu.Comp {
	t := gwu.NewTable()
	t.EnsureSize(5, 2)

	t.Add(gwu.NewLabel("Approximation"), 0, 0)
	self.info_approx = gwu.NewLabel("")
	t.Add(self.info_approx, 0, 1)

	t.Add(gwu.NewLabel("Degree"), 1, 0)
	self.info_degree = gwu.NewLabel("")
	t.Add(self.info_degree, 1, 1)

	t.Add(gwu.NewLabel("Iteration"), 2, 0)
	self.info_iter = gwu.NewLabel("")
	t.Add(self.info_iter, 2, 1)

	t.Add(gwu.NewLabel("Max Error"), 3, 0)
	self.info_error = gwu.NewLabel("")
	t.Add(self.info_error, 3, 1)

	t.Add(gwu.NewLabel("Optimality"), 4, 0)
	self.info_optimality = gwu.NewLabel("")
	t.Add(self.info_optimality, 4, 1)
	return t
}

func (self *ApproxGUI) buildErrorTable() gwu.Comp {
	self.error_table = gwu.NewTable()
	self.error_table.SetCellPadding(4)
	self.error_table.Style().SetBorder2(1, gwu.BRD_STYLE_SOLID, gwu.CLR_BLACK)
	return self.error_table
}