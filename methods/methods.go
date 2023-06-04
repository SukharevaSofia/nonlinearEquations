package methods

import (
	"fmt"
	_ "github.com/wcharczuk/go-chart/v2"
	chart "github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"math"
	"nonlinearEquations/utils"
	"os"
	"strconv"
)

const dx = 0.000001

// считает производную
func derive(f func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return (f(x+dx) - f(x)) / dx
	}
}

func CheckNumberOfRoots(f func(x float64) float64, a, b float64, accuracy int) {
	step := (b - a) / math.Pow(10, float64(accuracy))
	cnt := 0
	for a <= b {
		if f(a)*f(a+step) <= 0 {
			cnt++
		}
		a += step
	}

	switch cnt {
	case 1:
		fmt.Println("На отрезке 1 корень")
	case 0:
		fmt.Println("На отрезке нет корней")
		os.Exit(1)
	default:
		fmt.Println("На отрезке более одного корня")
		os.Exit(1)
	}

}

func checkConvergence(f func(x float64) float64, a, b float64, accuracy int) bool {
	d := derive(f)
	if !(f(a)*f(b) <= 0) {
		fmt.Printf("Не выполняется достаточное условие единственности корня на отрезке [%.3f,%.3f]", a, b)
		return false
	}
	sign := d(a) >= 0
	cnt := 0
	step := (b - a) / math.Pow(10, float64(accuracy))
	for a <= b {
		if (d(a) >= 0) != sign {
			cnt++
			sign = !sign
		}
		a += step
	}
	return true
}

func ShowGraph(fileName string, f func(x float64) float64, a, b, x float64) {
	cntOfRepeat := 1000
	X := []float64{x}
	y := f(x)
	Y := []float64{y}
	min := a - math.Abs(b-a)
	max := b + 2*math.Abs(b-a)
	listX := generateX(min, max, cntOfRepeat)
	listY := generateY(listX, f)
	legend := "X: " + strconv.FormatFloat(x, 'e', 3, 64) +
		"; Y: " + strconv.FormatFloat(y, 'e', 3, 64)
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "X",
			GridLines: []chart.GridLine{{false, chart.Style{Hidden: false, StrokeWidth: 3,
				StrokeColor: drawing.ColorFromAlphaMixedRGBA(0, 0, 0, 255)}, 0.0},
				{false, chart.Style{Hidden: false, StrokeWidth: 3,
					StrokeColor: drawing.ColorFromAlphaMixedRGBA(255, 255, 0, 255)}, a},
				{false, chart.Style{Hidden: false, StrokeWidth: 3,
					StrokeColor: drawing.ColorFromAlphaMixedRGBA(255, 255, 0, 255)}, b}},
		},
		YAxis: chart.YAxis{
			Name: "Y",
			GridLines: []chart.GridLine{
				{false, chart.Style{Hidden: false, StrokeWidth: 3,
					StrokeColor: drawing.ColorFromAlphaMixedRGBA(0, 0, 0, 255)}, 0.0}},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    1,
				},
				Name:    fileName,
				XValues: listX,
				YValues: listY,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 3,
					DotColor: chart.ColorRed,
				},
				Name:    legend,
				XValues: X,
				YValues: Y,
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	picture, _ := os.Create(fileName)
	graph.Render(chart.PNG, picture)
	picture.Close()
}

func ShowSysGraph(xGiven, yGiven, xAns, yAns float64, eqnSystem utils.SysOfEqn) {
	cntOfRepeat := 1000
	min := xGiven - 3
	max := xGiven + 2
	listX1, listX2 := generateSysX(min, max, cntOfRepeat, eqnSystem.FirstYfromX)
	if len(listX2) == 0 {
		listX2 = listX1
	}
	listY1top, listY1bottom := generateSysY(listX1, eqnSystem.FirstYfromX)
	if len(listY1bottom) == 0 {
		listY1bottom = listY1top
	}
	listY2, _ := generateSysY(listX2, eqnSystem.SecondYfromX)
	Xanswers := []float64{xAns}
	Yanswers := []float64{yAns}
	Xgiven := []float64{xGiven}
	Ygiven := []float64{yGiven}
	givenValuesString := fmt.Sprintf("Начальное приближение: %.3f; %.3f", xGiven, yGiven)
	legend := "X: " + strconv.FormatFloat(xAns, 'e', 3, 64) +
		"; Y: " + strconv.FormatFloat(yAns, 'e', 3, 64)
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "X",
			GridLines: []chart.GridLine{{false, chart.Style{Hidden: false, StrokeWidth: 3,
				StrokeColor: drawing.ColorFromAlphaMixedRGBA(0, 0, 0, 255)}, 0.0},
			},
		},
		YAxis: chart.YAxis{
			Name: "Y",
			GridLines: []chart.GridLine{
				{false, chart.Style{Hidden: false, StrokeWidth: 3,
					StrokeColor: drawing.ColorFromAlphaMixedRGBA(0, 0, 0, 255)}, 0.0}},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    1,
					DotColor:    chart.ColorBlue,
				},
				Name:    "f",
				XValues: listX1,
				YValues: listY1top,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    1,
					DotColor:    chart.ColorBlue,
				},
				XValues: listX1,
				YValues: listY1bottom,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    1,
				},
				Name:    "g",
				XValues: listX2,
				YValues: listY2,
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 3,
					DotColor: chart.ColorCyan,
				},
				Name:    legend,
				XValues: Xanswers,
				YValues: Yanswers,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 3,
					DotColor: chart.ColorOrange,
				},
				Name:    givenValuesString,
				XValues: Xgiven,
				YValues: Ygiven,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 1,
					DotColor: chart.ColorBlack,
				},
				XValues: []float64{0},
				YValues: []float64{-5},
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	picture, _ := os.Create("system.png")
	graph.Render(chart.PNG, picture)
	picture.Close()
}

func generateX(min, max float64, cntOfRepeat int) []float64 {
	var list []float64
	step := (max - min) / float64(cntOfRepeat)
	value := min
	for value <= max {
		list = append(list, value)
		value += step
	}
	return list
}
func generateY(list []float64, f func(val float64) float64) []float64 {
	var listY []float64
	for _, value := range list {
		listY = append(listY, f(value))
	}
	return listY
}
func generateSysX(min, max float64, cntOfRepeat int, f func(val float64) (float64, float64)) ([]float64, []float64) {
	var list1 []float64
	var list2 []float64
	step := (max - min) / float64(cntOfRepeat)
	value := min
	for value <= max {
		a, b := f(value)
		if !math.IsNaN(a) {
			list1 = append(list1, value)
		}
		if !math.IsNaN(b) {
			list2 = append(list2, value)
		}
		value += step
	}
	return list1, list2
}
func generateSysY(list []float64, f func(val float64) (float64, float64)) ([]float64, []float64) {
	var listY1 []float64
	var listY2 []float64
	for _, value := range list {
		a, b := f(value)
		if !math.IsNaN(a) {
			listY1 = append(listY1, a)
		} else {
			listY1 = append(listY1, -100)
		}
		if (!math.IsNaN(b)) && (b != math.Inf(1)) {
			listY2 = append(listY2, b)
		} else if math.IsNaN(b) {
			listY2 = append(listY2, -100)
		}

	}
	return listY1, listY2
}
