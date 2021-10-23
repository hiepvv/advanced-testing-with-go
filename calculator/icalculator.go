package calculator

type ICalculator interface {
	Plus(a, b float64) float64
	Minus(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) float64
}
