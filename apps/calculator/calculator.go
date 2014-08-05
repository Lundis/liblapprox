package main

import (
	//"code.google.com/p/liblundis/lmath/algebra"
	. "code.google.com/p/liblundis/lmath/calculator"
	"code.google.com/p/gowut/gwu"
	"fmt"
	"os"
)

type SessHandler struct{}

func (h SessHandler) Created(s gwu.Session) {
	fmt.Println("SESSION created:", s.Id())
	buildGUI(s)
}

func (h SessHandler) Removed(s gwu.Session) {
	fmt.Println("SESSION removed:", s.Id())
}


func buildGUI(s gwu.Session) {
	win := gwu.NewWindow("main", "Calculator")
	p := gwu.NewPanel()

	t := gwu.NewTabPanel()
	t.Style().SetSizePx(800, 600)
	calc := BuildCalculatorGUI()
	t.AddString("Calculator", calc)
	approx := BuildApproximatorGUI()
	t.AddString("Approximator", approx)

	p.Add(t)
	win.Add(p)
	s.AddWin(win)
}


func main() {
	server := gwu.NewServer("luncalc", "localhost:8082")
	server.SetText("Lundis Calculator")
	server.AddSessCreatorName("main", "Calculator")
	server.AddSHandler(SessHandler{})
	startConsoleReader()
	server.Start("")

}

func startConsoleReader() {
	fmt.Println("Type 'r' to restart, 'e' to exit.")
	go func() {
		var cmd string
		for {
			fmt.Scanf("%s", &cmd)
			switch cmd {
			case "r": // restart
				os.Exit(1)
			case "e": // exit
				os.Exit(0)
			}
		}
	}()
}