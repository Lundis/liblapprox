package calculator

import (
	"code.google.com/p/gowut/gwu"
)

func BuildApproximatorGUI() gwu.Comp {
	p := gwu.NewPanel()
	title := gwu.NewLabel("Approximator!")
	p.Add(title)
	return p
}