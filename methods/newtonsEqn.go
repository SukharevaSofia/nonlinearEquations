package methods

import "math"

func NewtonsEqn(f func(x float64) float64, a, b float64, acc int) (float64, float64, int) {
	result := checkConvergence(f, a, b, acc)
	if !result {
		return 0, 0, 0
	}
	x := getFirstApproximation(f, a, b)
	x0 := x
	cntOfIterations := 0
	d := derive(f)
	accuracy := float64(acc)
	if math.Abs(f(x)) > math.Pow(10, -accuracy) &&
		math.Abs(x-x0) > math.Pow(10, -accuracy) &&
		math.Abs(f(x)/d(x)) > math.Pow(10, -accuracy) {
		return x, f(x), 0
	} else {
		for {
			x0 = x
			cntOfIterations++
			x = x0 - f(x0)/d(x)
			if !(math.Abs(f(x)) > math.Pow(10, -accuracy) &&
				math.Abs(x-x0) > math.Pow(10, -accuracy) &&
				math.Abs(f(x)/d(x)) > math.Pow(10, -accuracy)) {
				break
			}
		}
		ShowGraph("newton.png", f, a, b, x)
		return x, f(x), cntOfIterations
	}
}

func getFirstApproximation(f func(x float64) float64, a, b float64) float64 {
	secondD := derive(derive(f))
	if f(a)*secondD(a) > 0 {
		return a
	}
	return b
}
