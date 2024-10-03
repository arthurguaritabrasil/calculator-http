package models

import "math"

type Pow struct {
	Operation
}

func (p Pow) CalculateOp(a, b float32, op string) float32 {
	p.operation = op
	p.result = float32(math.Pow(float64(a), float64(b)))
	return p.result
}
