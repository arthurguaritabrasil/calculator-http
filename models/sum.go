package models

type Sum struct {
	Operation
}

func (s Sum) calculateOp(a, b float32, op string) float32 {

	s.operation = op
	s.result = a + b
	return s.result

}
