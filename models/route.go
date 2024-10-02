package models

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

var Routes = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"GET": {
		"/result": Runner,
	},
}

func Runner(w http.ResponseWriter, r *http.Request) {

	op := r.URL.Query().Get("op")

}
