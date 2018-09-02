package routes

import (
	"net/http"
)

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type EditItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Query    string `json:"query"`
}

type ItemListPrice struct {
	Items []Item `json:"item"`
	Total int    `json:"total"`
}

var itemsBunch ItemListPrice

func Api(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/edit", edit)
	mux.HandleFunc("/add", add)
	mux.HandleFunc("/delete", delete)
	return mux
}
