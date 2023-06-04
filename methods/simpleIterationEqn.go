package methods

import (
	"fmt"
	"math"
)

func SimpleIterEqn(f func(x float64) float64, a, b, q float64, accuracy int) (float64, float64, int) {
	lambda := getLambdaIter(f, a, b)
	phi := func(x float64) float64 { return x + lambda*f(x) }
	var x0 float64
	if f(a) > f(b) {
		x0 = a
	} else {
		x0 = b
	}
	x := x0
	cntOfIterations := 0
	for {
		cntOfIterations++
		x0 = x
		//x = x - f(x)/lambda
		x = phi(x0)
		if checkEnd(x, x0, float64(accuracy)) {
			break
		}
		if math.IsInf(x, 1) || math.IsNaN(x) || math.IsInf(x, -1) {
			fmt.Printf("На промежутке [%.3f, %.3f] метод простых иттераций расходится.", a, b)
			return 0, 0, 0
		}

	}
	ShowGraph("iterations.png", f, a, b, x)
	return x, f(x), cntOfIterations
}

func getLambdaIter(f func(x float64) float64, a, b float64) float64 {
	d := derive(f)
	return -1 / math.Max(math.Abs(d(a)), math.Abs(d(b)))
}

func checkEnd(x, x0, accuracy float64) bool {
	return math.Abs(x-x0) <= math.Pow(10, -accuracy)
}
