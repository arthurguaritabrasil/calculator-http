package models

import "math"

type Sqrt struct {
	Operation
}

func (s Sqrt) CalculateOp(a, b float32, op string) float32 {

	if a < 0 {
		return -1
	}

	s.operation = op
	s.result = float32(math.Pow(float64(a), 1/float64(b)))
	return s.result
}
