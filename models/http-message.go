package models

type HttpMessage struct {
	Code  int    `json: "code"`
	Error string `json: "error"`
}
