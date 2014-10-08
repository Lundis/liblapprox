package main

import (
	"code.google.com/p/gowut/gwu"
	"os"
	"fmt"
)

type ApproxBackend interface {

	// Returns a String representation of the approximation of specified 
	// degree (and if applicable, iteration)
	String(deg, iter int) string

	// Returns an url to the requested image
	ImageUrl(deg, iter, dimx, dimy int) string

	ErrorGraphUrl(deg, iter, dimx, dimy int) string

	Error(deg, iter int) float64

	// Returns (if applicable), an upper bound as to how close this approximation is to the optimal one for this degree
	Optimality(deg, iter int) float64 

	// Returns the number of iterations for a specific degree
	Iters(deg int) int

	BuildInfoTable() gwu.Comp

	UpdateInfoTable(deg, iter int)

	Filename(deg, iter, dimx, dimy int) string
}

type ApproxBackendImpl struct {
	id string
}

func (self *ApproxBackendImpl) Filename(deg, iter, dimx, dimy int) string {
	return fmt.Sprintf("%v_%v_%v_%v_%v.png", self.id, deg, iter, dimx, dimy)
}

func imageUrl(backend ApproxBackend, deg, iter, dimx, dimy int) (string, bool) {
	filename := backend.Filename(deg, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	return "img/" + filename, existsFile(full_path)
}

func errorGraphUrl(backend ApproxBackend, deg, iter, dimx, dimy int) (string, bool) {
	filename := "err_" + backend.Filename(deg, iter, dimx, dimy)
	full_path := ImageDir() + string(os.PathSeparator) + filename
	return "img/" + filename, existsFile(full_path)
}

func existsFile(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}