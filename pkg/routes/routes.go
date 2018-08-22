package routes

import (
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../web/app/"))
	mux.Handle("/", fs)
	Api(mux)
	return mux
}
