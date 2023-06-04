package utils

type SysOfEqn struct {
	First        func(x, y float64) float64
	Second       func(x, y float64) float64
	FirstYfromX  func(x float64) (float64, float64)
	SecondYfromX func(x float64) (float64, float64)
	AnswX, AnswY float64
}

type Matrix struct {
	First  []func(x, y float64) float64
	Second []func(x, y float64) float64
}
