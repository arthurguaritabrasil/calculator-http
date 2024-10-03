package utils

import (
	"net/http"
)

func ItsMethodPathValid(r *http.Request) bool {
	return r.Method == "GET" && r.URL.Path == "/result"
}

func ItsOperationValid(op string) bool {
	return Operations[op] != nil
}
