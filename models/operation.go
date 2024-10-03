package models

type Operation struct {
	a         float32
	b         float32
	result    float32 `json: "result"`
	operation string  `json: "op"`
}

type Operator interface {
	CalculateOp(a, b float32, op string) float32
}
