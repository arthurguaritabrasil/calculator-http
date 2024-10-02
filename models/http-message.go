package models

import (
	"encoding/json"
	"net/http"
)

type HttpMessage struct {
	Code  int
	Error string
}

var httpMessages = map[int]HttpMessage{
	404: HttpMessage{404, "not found"},
	500: HttpMessage{500, "something went wrong"},
}

func WriteResultError(w http.ResponseWriter, code int) {
	jon, _ := json.Marshal(httpMessages[code])
	w.Write(jon)
}
