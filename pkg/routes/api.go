package routes

import "net/http"

func Api(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/edit", edit)
	mux.HandleFunc("/add", edit)
	mux.HandleFunc("/delete", edit)
	return mux
}

func edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ..

	w.WriteHeader(http.StatusOK)
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ..

	w.WriteHeader(http.StatusOK)
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ..

	w.WriteHeader(http.StatusOK)
}
