package models

type Operation struct {
	a         float32
	b         float32
	result    float32
	operation string
}

type Operator interface {
	calculateOp(a, b float32, op string) float32
}
