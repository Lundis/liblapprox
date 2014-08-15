package main

import (
)

type ApproxBackend interface {

	// Returns a String representation of the approximation of specified 
	// degree (and if applicable, iteration)
	String(degree, iter int) string

	// Returns an url to the requested image
	ImageUrl(degree, iter, dimx, dimy int) string

	Error(degree, iter int) float64

	// Returns (if applicable), an upper bound as to how close this approximation is to the optimal one for this degree
	Optimality(degree, iter int) float64 

	// Returns the number of iterations for a specific degree
	Iters(degree int) int
}