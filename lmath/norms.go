package lmath

func LInfNorm1f1(f Func1to1, start, end float64) float64 {
	return FindMaxDiff(ZeroFunc(), f, start, end)
}

func LpNorm1f1(f Func1to1) float64 {
	return 0
}