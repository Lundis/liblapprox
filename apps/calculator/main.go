package main

import (
	//"code.google.com/p/liblundis/lmath/algebra"
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
	//t.Style().SetSizePx(1024, 600)
	
	approx := new(ApproxGUI)
	approx_comp := approx.BuildGUI()
	t.AddString("Approximator", approx_comp)

	calc := BuildCalculatorGUI()
	t.AddString("Calculator", calc)

	p.Add(t)
	win.Add(p)
	s.AddWin(win)
}


func main() {
	server := gwu.NewServer("luncalc", "localhost:8082")
	server.SetText("Lundis Calculator")
	server.AddSessCreatorName("main", "Approximator")
	server.AddSHandler(SessHandler{})

	img_dir := ImageDir()
	os.MkdirAll(img_dir, 0777)
	server.AddStaticDir("img", img_dir)
	defer os.RemoveAll(img_dir)

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