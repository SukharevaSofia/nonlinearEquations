package main

import (
	"fmt"
	"math"
	"nonlinearEquations/methods"
	"nonlinearEquations/utils"
)

// Решение нелинейных уравнений
// Метод хорд
// Метод Ньютона
// Метод простой итерации
// Метод Ньютона для системы
func main() {
	var equation func(x float64) float64
	eq1 := []float64{1, -1, 4}
	eq2 := []float64{2, -1}
	eq3 := []float64{0.5}
	fNumber, a, b, accuracy, q := inputValues(eq1, eq2, eq3)
	switch fNumber {
	case 1:
		equation = func(x float64) float64 {
			return eq1[0]*x*x*x + eq1[1]*x + eq1[2]
		}
	case 2:
		equation = func(x float64) float64 {
			return math.Exp(eq2[0]*x) + eq2[1]
		}
	case 3:
		equation = func(x float64) float64 {
			return math.Sin(eq3[0] * x)
		}

	}

	methods.CheckNumberOfRoots(equation, a, b, accuracy)
	newtX, newtY, newtIt := methods.NewtonsEqn(equation, a, b, accuracy)
	if newtIt != 0 {
		fmt.Println("Ньютон: ", newtX, newtY, newtIt)
	}
	simpX, simpY, simpit := methods.SimpleIterEqn(equation, a, b, q, accuracy)
	if simpit != 0 {
		fmt.Println("Итерации: ", simpX, simpY, simpit)
	}
	chorX, chorY, chorIt := methods.ChordEqn(equation, a, b, accuracy)
	if chorIt != 0 {
		fmt.Println("Хорды: ", chorX, chorY, chorIt)
	}

	s1 := [][]float64{{1, 1, 4}, {3, 0, 1}}
	s2 := [][]float64{{3, -1, 1}, {1, 1, 1}}
	sNumber, x0, y0 := eqnSysInput(s1, s2)

	var sysEq utils.SysOfEqn
	switch sNumber {
	case 1:
		sysEq.First = func(x, y float64) float64 {
			return s1[0][0]*x*x + s1[0][1]*y*y - s1[0][2]
		}
		sysEq.Second = func(x, y float64) float64 {
			return s1[1][0]*x*x + s1[1][1] - s1[1][2]*y
		}
		sysEq.FirstYfromX = func(x float64) (float64, float64) {
			return math.Sqrt(4 - x*x), -math.Sqrt(4 - x*x)
		}
		sysEq.SecondYfromX = func(x float64) (float64, float64) {
			return (s1[1][0]*x*x + s1[1][1]) / s1[1][2], math.Inf(1)
		}
		answ := 0.783212564406379
		if math.Abs(x0-answ) < math.Abs(x0+answ) {
			sysEq.AnswX = answ
		} else {
			sysEq.AnswX = -answ
		}
		sysEq.AnswY = 1.840265763132
	case 2:
		sysEq.First = func(x, y float64) float64 {
			return s2[0][0] + s2[0][1]*x*x - s2[0][2]*y
		}
		sysEq.Second = func(x, y float64) float64 {
			return s2[1][0]*x + s2[1][1] - s2[1][2]*y
		}
		sysEq.FirstYfromX = func(x float64) (float64, float64) {
			return (s2[0][0] + s2[0][1]*x*x) / s2[0][2], math.Inf(1)
		}
		sysEq.SecondYfromX = func(x float64) (float64, float64) {
			return (s2[1][0]*x + s2[1][1]) / s2[1][2], math.Inf(1)
		}
		answx1, answy1, answx2, answy2 := -2., -1., 1., 2.
		if math.Sqrt(math.Pow(x0-answx1, 2)+math.Pow(y0-answy1, 2)) <
			math.Sqrt(math.Pow(x0-answx2, 2)+math.Pow(y0-answy2, 2)) {
			sysEq.AnswX, sysEq.AnswY = answx1, answy1
		} else {
			sysEq.AnswX, sysEq.AnswY = answx2, answy2
		}
	}

	result, xSys, ySys, sysIt := methods.NewtonSys(sysEq, x0, y0, accuracy)
	fmt.Println(result)
	//answX, answY := sysEq.AnswX, sysEq.AnswY

	fmt.Printf("X: %.3f; Y: %.3f\nЗначения: %.3e; %.3e\n", xSys, ySys,
		sysEq.First(xSys, ySys), sysEq.Second(xSys, ySys))
	fmt.Println("Итераций: ", sysIt)

}
