package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adrianforsius/backend-challenge/product"
)

func main() {
	baskets := product.NewBasket()

	router := http.NewServeMux()
	router.Handle("/checkout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			prod, err := baskets.Get("1")
			if err != nil {
				http.Error(w, fmt.Sprintf("{\"error:\": \"failed %s\"}", err), http.StatusNotFound)
				return
			}

			data, err := json.Marshal(prod)
			if err != nil {
				http.Error(w, fmt.Sprintf("{\"error:\": \"failed %s\"}", err), http.StatusInternalServerError)
				return
			}
			w.Write(data)
			return
		case http.MethodPost:
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}))

	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}
