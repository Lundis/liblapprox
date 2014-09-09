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

    p := createPlot(title)

    for i := range funcs {
        f := plotter.NewFunction(funcs[i])
        f.Color = colors[i % len(colors)]
        f.Samples = w
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
    diff := max - min
    p.Y.Max = max + diff*0.05
    p.Y.Min = min - diff*0.05

    savePlot(p, filename, w, h)
}

func SaveSimpleGraphY(y_coords [][]float64, labels []string, title string, filename string, w, h int) {
    x_coords := make([][]float64, len(y_coords))
    // fill x_coords with 0, 1, 2, 3, ...
    for i := range x_coords {
        x_coords[i] = make([]float64, len(y_coords[i]))
        for j := range x_coords[i] {
            x_coords[i][j] = float64(j)
        }
    }
    SaveSimpleGraphXY(x_coords, y_coords, labels, title, filename, w, h)
}

func SaveSimpleGraphXY(x_coords, y_coords [][]float64, labels []string, title string, filename string, w, h int) {
    if  len(x_coords) != len(y_coords) {
        panic("SaveSimpleGraph: x and y must have equal amount of elements.")
    }
    if len(x_coords) != len(labels) {
        panic("SaveSimpleGraph: all plots must have labels.")
    }
    if len(x_coords) == 0 {
        panic("SaveSimpleGraph: no coordinates given.")
    }
    p := createPlot(title)

    for i, x_vals := range x_coords {
        y_vals := y_coords[i]
        f, err := plotter.NewLine(convertCoords(x_vals, y_vals))
        if err != nil {
            panic(err)
        }
        f.Color = colors[i % len(colors)]
        p.Add(f)
        p.Legend.Add(labels[i], f)
    }

    savePlot(p, filename, w, h)
}

func createPlot(title string) *plot.Plot {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = title
    p.X.Label.Text = "x"
    p.Y.Label.Text = "y"
    return p
}

func savePlot(p *plot.Plot, filename string, w, h int) {
    if err := p.Save(float64(w)/96, float64(h)/96, filename); err != nil {
        panic(err)
    }
}

func convertCoords(x, y []float64) (plotter.XYs) {
    pts := make(plotter.XYs, len(x))
    for i := range x {
        pts[i].X = x[i]
        pts[i].Y = y[i]
    }
    return pts
}