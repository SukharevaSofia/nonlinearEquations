package methods

import (
	"math"
)

func ChordEqn(f func(x float64) float64, a, b float64, acc int) (float64, float64, int) {
	a0, b0 := a, b
	accuracy := float64(acc)
	x := a - f(a)*(b-a)/(f(b)-f(a))
	x0 := x
	cntOfIterations := 0
	for {
		cntOfIterations++
		x0 = x
		x = a - f(a)*(b-a)/(f(b)-f(a))
		if f(a)*f(x) < 0 {
			a, b = a, x
		} else {
			a, b = x, b
		}
		if math.Abs(f(x)) <= math.Pow(10, -accuracy) ||
			(math.Abs(x-x0) <= math.Pow(10, -accuracy) && (cntOfIterations != 1)) {
			break
		}
	}
	ShowGraph("chords.png", f, a0, b0, x)
	return x, f(x), cntOfIterations
}
