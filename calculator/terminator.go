package calculator

type Terminator struct{}

func (c *Terminator) Plus(a, b float64) float64 {
	return a + b
}

func (c *Terminator) Minus(a, b float64) float64 {
	return a - b
}

func (c *Terminator) Multiply(a, b float64) float64 {
	return a * b
}

func (c *Terminator) Divide(a, b float64) float64 {
	return a / b
}
