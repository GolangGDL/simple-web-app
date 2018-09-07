package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Id       string `json:"id"`
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var elementEdit Item
		err := json.Unmarshal(bodyBytes, &elementEdit)
		if err != nil {
			log.Println("error:", err)
		}
		itemsBunch.Total = 0
		for index, element := range itemsBunch.Items {
			if element.Name == elementEdit.Id {
				itemsBunch.Items[index].Name = elementEdit.Name
				itemsBunch.Items[index].Price = elementEdit.Price
				itemsBunch.Items[index].Quantity = elementEdit.Quantity
			}
			totalProduct := productTotalPrice(itemsBunch.Items[index].Price, itemsBunch.Items[index].Quantity)
			itemsBunch.Total = newTotal(itemsBunch.Total, totalProduct)
		}
		elementEdit.Id = ""
		if err := json.NewEncoder(w).Encode(itemsBunch); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var element Item
		err := json.Unmarshal(bodyBytes, &element)
		if err != nil {
			log.Println("error:", err)
		}
		var indexDelete int
		for i, item := range itemsBunch.Items {
			if item.Name == element.Id {
				indexDelete = i
			}
		}
		var emptyItem Item
		copy(itemsBunch.Items[indexDelete:], itemsBunch.Items[indexDelete+1:])
		itemsBunch.Items[len(itemsBunch.Items)-1] = emptyItem
		itemsBunch.Items = itemsBunch.Items[:len(itemsBunch.Items)-1]
		itemsBunch.Total = 0
		for index := range itemsBunch.Items {
			totalProduct := productTotalPrice(itemsBunch.Items[index].Price, itemsBunch.Items[index].Quantity)
			itemsBunch.Total = newTotal(itemsBunch.Total, totalProduct)
		}
		if err := json.NewEncoder(w).Encode(itemsBunch); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func newTotal(total, newElement int) int {
	return total + newElement
}

func productTotalPrice(productPrice, quantity int) int {
	return productPrice * quantity
}
