package main

import (
	"bufio"
	"fmt"
	"math"
	"nonlinearEquations/utils"
	"os"
	"strconv"
	"strings"
)

func inputValues(eq1, eq2, eq3 []float64) (int, float64, float64, int, float64) {
	fmt.Println(utils.INFO)
	var acc int
	var a, b, q float64
	var left, right, accuracy, lipschitzCoefficient string
	fmt.Printf("Выберите функцию для интегрирования. Введите 1, 2 или 3\n"+
		"1: %.3fx^3 + (%.3fx) + (%.3f)\n"+
		"2: e^%.3fx + (%.3f)\n"+
		"3: sin(%.3fx)",
		eq1[0], eq1[1], eq1[2],
		eq2[0], eq2[1],
		eq3[0])
	fmt.Println(utils.REQUEST_FUNC)

	fNumber := 0
	for {
		fmt.Scan(&fNumber)
		if (fNumber == 1) || (fNumber == 2) || (fNumber == 3) {
			break
		}
		fmt.Println(utils.REQUEST_NUM)
		fmt.Print(utils.REQUEST_FUNC)
	}
	var inputString string
	fmt.Println(utils.CHOOSE_INPUT)
	for {
		fmt.Scan(&inputString)
		if inputString == "T" || inputString == "t" || inputString == "F" || inputString == "f" {
			break
		}
		fmt.Println(utils.CHOOSE_INPUT)
	}
	if inputString == "F" || inputString == "f" {
		f, _ := os.Open("data1.txt")
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		limits := strings.Split(scanner.Text(), " ")
		left, right, accuracy, lipschitzCoefficient = limits[0], limits[1], limits[2], limits[3]
		a, _ = strconv.ParseFloat(left, 64)
		b, _ = strconv.ParseFloat(right, 64)
		if a > b {
			fmt.Print(utils.INPUT_ERR)
			os.Exit(1)
		}

	} else {
		fmt.Print(utils.REQUEST_LEFT)
		for {
			fmt.Scan(&left)
			if _, err := strconv.ParseFloat(left, 64); err == nil {
				break
			}
			fmt.Print(utils.INPUT_ERR)
		}
		fmt.Print(utils.REQUEST_RIGHT)
		for {
			fmt.Scan(&right)
			if a, err := strconv.ParseFloat(right, 64); err == nil {
				if b, _ = strconv.ParseFloat(left, 64); a > b {
					break
				}

			}
			fmt.Print(utils.INPUT_ERR)
		}

		fmt.Print(utils.REQUEST_ACCURACY)
		for {
			fmt.Scan(&accuracy)
			if _, err := strconv.ParseFloat(accuracy, 64); err == nil {
				break
			}
			fmt.Print(utils.INPUT_ERR)
		}
		fmt.Print(utils.REQUEST_Q)
		for {
			fmt.Scan(&lipschitzCoefficient)
			if q, err := strconv.ParseFloat(lipschitzCoefficient, 64); err == nil {
				if q > 0 && q < 1 {
					break
				}
			}
		}
	}
	c, _ := strconv.ParseFloat(accuracy, 64)
	for math.Mod(c, 1) != 0 {
		acc++
		c *= 10
	}

	a, _ = strconv.ParseFloat(left, 64)
	b, _ = strconv.ParseFloat(right, 64)
	q, _ = strconv.ParseFloat(lipschitzCoefficient, 64)
	return fNumber, a, b, acc, q
}

func eqnSysInput(s1, s2 [][]float64) (int, float64, float64) {
	fmt.Printf("Выберите Систему уравнений. Введите 1 или 2 \n"+
		"1: \n"+
		" %.3fx^2 + (%.3fy^2) = (%.3f)\n"+
		" %.3fx^2 + (%.3f) = %.3fy\n",
		s1[0][0], s1[0][1], s1[0][2], s1[1][0], s1[1][1], s1[1][2])
	fmt.Printf("2:\n"+
		" %.3f + (%.3fx^2) = %.3fy\n"+
		" %.3fx + (%.3f) = %.3fy\n",
		s2[0][0], s2[0][1], s2[0][2], s2[1][0], s2[1][1], s2[1][2])

	fNumber := 0
	for {
		fmt.Scan(&fNumber)
		if (fNumber == 1) || (fNumber == 2) {
			break
		}
		fmt.Print(utils.REQUEST_SYS)
	}

	var inputString string
	fmt.Println(utils.CHOOSE_INPUT)
	for {
		fmt.Scan(&inputString)
		if inputString == "T" || inputString == "t" || inputString == "F" || inputString == "f" {
			break
		}
		fmt.Println(utils.CHOOSE_INPUT)
	}
	var x0s, y0s string
	var x0, y0 float64
	if inputString == "F" || inputString == "f" {
		f, _ := os.Open("data2.txt")
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		limits := strings.Split(scanner.Text(), " ")
		x0s, y0s = limits[0], limits[1]
	} else {
		fmt.Print(utils.REQUEST_X0)
		for {
			fmt.Scan(&x0s)
			if _, err := strconv.ParseFloat(x0s, 64); err == nil {
				break
			}
			fmt.Print(utils.INPUT_ERR)
		}
		fmt.Print(utils.REQUEST_Y0)
		for {
			fmt.Scan(&y0s)
			if _, err := strconv.ParseFloat(y0s, 64); err == nil {
				break
			}
			fmt.Print(utils.INPUT_ERR)
		}
	}
	x0, _ = strconv.ParseFloat(x0s, 64)
	y0, _ = strconv.ParseFloat(y0s, 64)
	return fNumber, x0, y0
}
