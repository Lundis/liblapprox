package lmath

func LInfNorm1f1(f Function, start, end float64) (float64, float64) {
	return FindMaxDiff(ZeroFunc(), f, start, end)
}

// Not Implemented...
func LpNorm1f1(f Function) float64 {
	return 0
}