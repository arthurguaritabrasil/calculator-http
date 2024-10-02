package server

import (
	"log"
	"net/http"
	"time"
)

type myHandler struct{}

func (my myHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {

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
