package server

import (
	"backend/models"
	"log"
	"net/http"
	"time"
)

type myHandler struct{}

func (my myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	models.Routes[r.Method][r.URL.Path](w, r)
}

func StartServer() {
	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      myHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
