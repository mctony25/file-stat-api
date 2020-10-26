package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

type MiddleWare struct {
	Router *mux.Router
}

func (mid *MiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	}

	if r.Method == "OPTIONS" {
		return
	}

	mid.Router.ServeHTTP(w, r)
}
