package methods

import (
	"fmt"
	"math"
)

func SimpleIterEqn(f func(x float64) float64, a, b, q float64, accuracy int) (float64, float64, int) {
	lambda := getLambdaIter(f, a, b)
	d := derive(f)
	phi := func(x float64) float64 { return x + lambda*f(x) }
	phiD := func(x float64) float64 { return 1 + lambda*d(x) }
	fmt.Println("PHI: ", phiD(a), phiD(b))
	var x0 float64

	x := b
	cntOfIterations := 0
	for {
		cntOfIterations++
		x0 = x
		//x = x - f(x)/lambda
		x = phi(x0)
		if math.Abs(f(x)) <= math.Pow(10, -float64(accuracy)) {
			break
		}
		if cntOfIterations > 500 || math.IsInf(x, 1) || math.IsNaN(x) || math.IsInf(x, -1) {
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
