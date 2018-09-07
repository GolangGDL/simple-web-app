package routes

import (
	"net/http"
)

func Api(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/edit", edit)
	mux.HandleFunc("/add", add)
	mux.HandleFunc("/delete", delete)
	return mux
}
