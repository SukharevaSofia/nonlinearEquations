package methods

import (
	"math"
	"nonlinearEquations/utils"
)

func NewtonSys(equations utils.SysOfEqn, x0, y0 float64, acc int) (string, float64, float64, int) {

	accuracy := float64(acc)
	xGiven, yGiven := x0, y0
	f := equations.First
	g := equations.Second
	jacobiMatrix := getJacobiMatrix(f, g)
	listOfFunc := []func(x, y float64) float64{equations.First, equations.Second}
	dxNew, dyNew := getNewApproximation(jacobiMatrix, listOfFunc, x0, y0)
	cntOfIterations := 1
	for math.Abs(dxNew) > math.Pow(10, -accuracy) || math.Abs(dyNew) > math.Pow(10, -accuracy) {
		x0 += dxNew
		y0 += dyNew
		newApproximationX, newApproximationY := getNewApproximation(jacobiMatrix, listOfFunc, x0, y0)
		cntOfIterations++
		dxNew = newApproximationX
		dyNew = newApproximationY
	}
	x := x0 + dxNew
	y := y0 + dyNew
	result := checkAnswer(x, y, equations.AnswX, equations.AnswY, math.Pow(10, -accuracy))

	ShowSysGraph(xGiven, yGiven, x, y, equations)
	return result, x, y, cntOfIterations
}

func getJacobiMatrix(f, g func(x float64, y float64) float64) utils.Matrix {
	var matrix utils.Matrix
	row1 := []func(x, y float64) float64{deriveX(f), deriveY(f)}
	row2 := []func(x, y float64) float64{deriveX(g), deriveY(g)}
	matrix.First = row1
	matrix.Second = row2
	return matrix
}
func getNewApproximation(jacobiMatrix utils.Matrix, functions []func(x float64, y float64) float64, x0, y0 float64) (float64, float64) {
	m := jacobiMatrix
	dxNew := -m.First[1](x0, y0) / m.First[0](x0, y0)
	freeTerm := -functions[0](x0, y0) / m.First[0](x0, y0)
	dyNew := (-functions[1](x0, y0) - freeTerm*m.Second[0](x0, y0)) / ((m.Second[0](x0, y0) * dxNew) + m.Second[1](x0, y0))
	dxNew = freeTerm + dxNew*dyNew
	return dxNew, dyNew
}

const dX = 0.0001
const dY = 0.0001

func deriveX(f func(x, y float64) float64) func(x, y float64) float64 {
	return func(x, y float64) float64 { return (f(x+dX, y) - f(x, y)) / dX }
}

func deriveY(f func(x, y float64) float64) func(x, y float64) float64 {
	return func(x, y float64) float64 { return (f(x, y+dY) - f(x, y)) / dY }
}

func checkAnswer(x, y, answ1, answ2, accuracy float64) string {
	if (math.Abs(math.Abs(x)-answ1) <= accuracy) && (math.Abs(y-answ2) <= accuracy) {
		return utils.GOOD_SOLUTION
	} else {
		return utils.BAD_SOLUTION
	}
}
