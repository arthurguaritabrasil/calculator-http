package utils

import (
	"backend/models"
	"net/http"
)

func ItsMethodPathValid(r *http.Request) bool {
	return models.Routes[r.Method][r.URL.Path] != nil
}

func ItsOperationValid(op string) bool {
	return models.Operations[op] != nil
}
