package plot

import (
    "image/color"
    "math"
	"code.google.com/p/liblundis/lmath"
    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
)

var colors []color.Color = []color.Color{color.RGBA{A:255}, color.RGBA{R: 255, A:255}, color.RGBA{G: 255, A:255}, color.RGBA{R: 255, G: 255, A:255}, color.RGBA{R: 255, B: 255, A:255}, color.RGBA{G: 255, B: 255, A:255}}

func SaveSimpleGraph(funcs []lmath.Function, labels []string, start, end float64, title string, filename string, w, h int) {
    if len(funcs) != len(labels) {
        panic("SaveSimpleGraph: all functions must have labels")
    }
    if len(funcs) == 0 {
        panic("SaveSimpleGraph: empty function slice")
    }
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = title
    p.X.Label.Text = "x"
    p.Y.Label.Text = "y"

    for i := range funcs {
        f := plotter.NewFunction(funcs[i])
        f.Color = colors[i % len(colors)]
        p.Add(f)
        p.Legend.Add(labels[i], f)
    }

    p.X.Min = start
    p.X.Max = end

    min := funcs[0](start)
    max := min
    for i := range funcs {
        f_min, _ := lmath.Min(funcs[i], 100, start, end)
        min = math.Min(min, f_min)
        f_max, _ := lmath.Max(funcs[i], 100, start, end)
        max = math.Max(max, f_max)
    }
    p.Y.Max = max
    p.Y.Min = min

    if err := p.Save(float64(w)/96, float64(h)/96, filename); err != nil {
        panic(err)
    }
}