package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adrianforsius/backend-challenge/product"
)

type NewBasketResp struct {
	ID string `json:"id"`
}

func main() {
	baskets := product.NewBasket()

	router := http.NewServeMux()
	router.Handle("/checkout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("state: %+v\n", baskets)
		switch r.Method {
		case http.MethodGet:
			vars, ok := r.URL.Query()["basket_id"]
			if !ok {
				http.Error(w, "{\"error:\": \"missing basket id\"}", http.StatusBadRequest)
				return
			}

			prod, err := baskets.Get(vars[0])
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
			products := make([]product.Product, 0)
			err := json.NewDecoder(r.Body).Decode(&products)
			if err != nil {
				http.Error(w, fmt.Sprintf("{\"error:\": \"failed %s\"}", err), http.StatusNotFound)
			}

			id := baskets.Add(products)
			resp, err := json.Marshal(NewBasketResp{
				ID: id,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("{\"error:\": \"failed %s\"}", err), http.StatusInternalServerError)
				return
			}
			w.Write(resp)
			return
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}))

	err := http.ListenAndServe(":8082", router)
	log.Fatal(err)
}
