package plot

import (
	"os"
	"code.google.com/p/liblundis/lmath"
	"fmt"
)

// WritePlotData writes a table of function values to a file
// funcs: a map with the functions
// indices: The indices of the maps to be written
// n: amount of data points
func WritePlotData(funcs map[int] lmath.Func1to1, indices []int, filename string, n int, start, end float64) {
    file, _ := os.Create(filename)
    fmt.Fprintf(file, "x    \t")
    for _, i := range indices {
        fmt.Fprintf(file, "%d       ", i)
    }
    fmt.Fprintf(file, "\n")
    
    for i := 0; i <= n; i++ {
        x := start + (end - start) * (float64(i)/float64(n))
        fmt.Fprintf(file, "%+.3f\t", x)
        for _, v := range indices {
            fmt.Fprintf(file, "%+.4f ", funcs[v](x))
        }
        fmt.Fprintf(file, "\n")
    }
}