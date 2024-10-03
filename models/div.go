package models

type Div struct {
	Operation
}

func (d Div) CalculateOp(a, b float32, op string) float32 {
	d.operation = op
	d.result = a / b
	return d.result
}
