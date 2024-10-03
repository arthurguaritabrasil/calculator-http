package models

type Sub struct {
	Operation
}

func (s Sub) CalculateOp(a, b float32, op string) float32 {
	s.operation = op
	s.result = a - b
	return s.result
}
