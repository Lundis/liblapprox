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
	win := gwu.NewWindow("main", "Approximator")
	
	approx := new(ApproxGUI)
	approx_comp := approx.BuildGUI()
	win.Add(approx_comp)
	s.AddWin(win)
}


func main() {
	server := gwu.NewServer("luncalc", "localhost:8082")
	server.SetText("Lundis Approximator")
	server.AddSessCreatorName("main", "Approximator")
	server.AddSHandler(SessHandler{})

	img_dir := ImageDir()
	os.MkdirAll(img_dir, 0777)
	server.AddStaticDir("img", img_dir)
	defer os.RemoveAll(img_dir)

	server.Start("")

}