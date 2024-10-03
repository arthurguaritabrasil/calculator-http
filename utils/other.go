package utils

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var httpMessages = map[int]models.HttpMessage{
	404: models.HttpMessage{404, "not found"},
	500: models.HttpMessage{500, "something went wrong"},
}

var Operations = map[string]models.Operator{
	"+":   models.Sum{},
	"sum": models.Sum{},

	"-":   models.Sub{},
	"sub": models.Sub{},

	"*":   models.Mul{},
	"mul": models.Mul{},

	"/":   models.Div{},
	"div": models.Div{},

	"^":   models.Pow{},
	"pow": models.Pow{},

	"~":   models.Sqrt{},
	"rot": models.Sqrt{},
}

func Runner(w http.ResponseWriter, r *http.Request) {

	op := r.URL.Query().Get("op")

	sepOperation := SeparateString(op)

	a, _ := strconv.ParseFloat(sepOperation[0], 32)
	b, _ := strconv.ParseFloat(sepOperation[1], 32)
	operator := sepOperation[2]

	if !ItsOperationValid(operator) {
		WriteResultError(w, 500)
		return
	}

	resulted := Operations[operator].CalculateOp(float32(a), float32(b), operator)

	str := strconv.FormatFloat(float64(resulted), 'f', 2, 32)

	WriteResultSucess(w, str, float32(a), float32(b), operator)
}

func WriteResultError(w http.ResponseWriter, code int) {
	jon, _ := json.Marshal(httpMessages[code])
	w.Write(jon)
}

func WriteResultSucess(res http.ResponseWriter, resulted string, a, b float32, op string) {

	expression := fmt.Sprintf("%.2f %s %.2f = %s", a, op, b, resulted)

	var r = map[string]string{
		"result":    resulted,
		"operation": expression,
	}

	jon, _ := json.Marshal(r)

	res.Write(jon)
}

func SeparateString(op string) []string {

	var sep []string

	for key, _ := range Operations {
		for strings.Contains(op, key) {
			sep = strings.Split(op, key)
			sep = append(sep, key)
			return sep
		}
	}
	return nil

}
