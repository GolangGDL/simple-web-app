package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    int     `json:"price"`
	Id       float32 `json:"id"`
}

type ItemListPrice struct {
	Items []Item `json:"item"`
	Total int    `json:"total"`
}

var itemsBunch ItemListPrice

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var element Item
		err := json.Unmarshal(bodyBytes, &element)
		if err != nil {
			log.Println("error:", err)
		}
		itemsBunch.Items = append(itemsBunch.Items, element)
		totalProduct := productTotalPrice(element.Price, element.Quantity)
		itemsBunch.Total = newTotal(itemsBunch.Total, totalProduct)
		if err := json.NewEncoder(w).Encode(itemsBunch); err != nil {
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

func newTotal(total, newElement int) int {
	return total + newElement
}

func productTotalPrice(productPrice, quantity int) int {
	return productPrice * quantity
}
