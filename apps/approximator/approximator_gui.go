package main

import (
	"code.google.com/p/liblundis/lmath/util/cont"
	"code.google.com/p/gowut/gwu"
	"fmt"
)

type ApproxGUI struct {
	function_box         gwu.ListBox
	function             cont.Function
	custom_func          gwu.TextBox
	interval_box         gwu.TextBox
	ival_start, ival_end float64

	degrees_box       gwu.TextBox
	degrees         []int
	current_deg       int
	current_iter      int
	discrete_points []float64

	approx_box   gwu.ListBox
	accuracy_box gwu.TextBox
	accuracy     float64

	result_container  gwu.Panel
	degree_buttons    gwu.Panel
	iter_buttons      gwu.Panel

	dimx_box    gwu.TextBox
	dimy_box    gwu.TextBox
	dimx        int
	dimy        int
	image       gwu.Image
	error_image gwu.Image
	error_table gwu.Table

	err               error
	backend           ApproxBackend

	message         gwu.Label
	info_box        gwu.Panel
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
	self.updateErrorTable()
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
	self.error_image.SetUrl("")
	if self.backend != nil && self.err == nil {
		self.image.SetUrl(self.backend.ImageUrl(self.current_deg, self.current_iter, self.dimx, self.dimy))
		self.error_image.SetUrl(self.backend.ErrorGraphUrl(self.current_deg, self.current_iter, self.dimx, self.dimy))

	}
}

func (self *ApproxGUI) updateErrorTable() {
	self.error_table.Clear()
	rows := len(self.degrees) + 1
	cols := 1
	for _, deg := range self.degrees {
		iters := self.backend.Iters(deg)
		if iters > cols {
			cols = iters
		}
	}
	cols += 1
	self.error_table.EnsureSize(rows, cols)

	self.error_table.Add(gwu.NewLabel("degs\\iters"), 0, 0)
	// column header
	for c := 1; c < cols; c++ {
		self.error_table.Add(gwu.NewLabel(fmt.Sprintf("%v", c-1)), 0, c)
	}

	for r := 1; r < rows; r++ {
		deg := self.degrees[r-1]
		// row header
		self.error_table.Add(gwu.NewLabel(fmt.Sprintf("%v", deg)), r, 0)
		c := 1
		for ; c < self.backend.Iters(deg) + 1; c++ {
			iter := c - 1
			l := gwu.NewLabel(fmt.Sprintf("%.6f", self.backend.Error(deg, iter)))
			self.error_table.Add(l, r, c)
		}
		for ; c < cols; c++ {
			self.error_table.Add(gwu.NewLabel("-"), r, c)
		}
	}
}

func (self *ApproxGUI) updateInfoTable() {
	if self.backend != nil && self.err == nil {
		self.backend.UpdateInfoTable(self.current_deg, self.current_iter)
	}
}


func (self *ApproxGUI) approximateMinimax() {
	fmt.Println("Approximating minimax...")
	self.backend = NewMinimaxApprox(self)
	self.info_box.Clear()
	self.info_box.Add(self.backend.BuildInfoTable())
	self.current_deg = self.degrees[0]
	self.current_iter = 0
}

func (self *ApproxGUI) approximateMinimaxDiscrete() {
	fmt.Println("Approximating discrete minimax...")
	self.backend = NewDiscreteMinimaxApprox(self)
	self.info_box.Clear()
	self.info_box.Add(self.backend.BuildInfoTable())
	self.current_deg = self.degrees[0]
	self.current_iter = 0
}