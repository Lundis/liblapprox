package calculator

import (
	//"code.google.com/p/liblundis/lmath/algebra"
	"code.google.com/p/gowut/gwu"
)

func BuildCalculatorGUI() gwu.Comp {
	p := gwu.NewPanel()
	title := gwu.NewLabel("Calculator!")
	p.Add(title)
	return p
}