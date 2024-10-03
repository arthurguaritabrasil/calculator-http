package models

type Mul struct {
	Operation
}

func (m Mul) CalculateOp(a, b float32, op string) float32 {
	m.operation = op
	m.result = a * b
	return m.result
}
