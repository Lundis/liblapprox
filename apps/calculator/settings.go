package main

import(
	"os"
)

func ImageDir() string {
	return os.TempDir() + string(os.PathSeparator) + "l_approximator"
}