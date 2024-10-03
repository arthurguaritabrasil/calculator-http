package server

import (
	"backend/utils"
	"log"
	"net/http"
	"time"
)

var Routes = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"GET": {
		"/result": utils.Runner,
	},
}

type myHandler struct{}

func (my myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !utils.ItsMethodPathValid(r) {
		utils.WriteResultError(w, 404)
		return
	}

	Routes[r.Method][r.URL.Path](w, r)
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
