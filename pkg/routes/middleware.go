package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		if err := json.NewEncoder(w).Encode(bodyBytes); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func edit(w http.ResponseWriter, r *http.Request) {

	// ..
}

func delete(w http.ResponseWriter, r *http.Request) {

	// ..

}
