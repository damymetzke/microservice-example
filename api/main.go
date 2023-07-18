package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func list_products(w http.ResponseWriter, req *http.Request) {
  response, err := http.Get("http://product-catalog/products/")
	if err != nil {
		// Handle the error appropriately
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// Handle the non-OK status code appropriately
		http.Error(w, fmt.Sprintf("Server returned non-OK status: %v", response.StatusCode), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// Handle the error appropriately
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func list_users(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, `{"message":"Hello Users!"}`)
}

func list_orders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, `{"message":"Hello Orders!"}`)
}

func create_order(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/products", list_products).Methods("Get")
	mux.HandleFunc("/users", list_users).Methods("Get")
	mux.HandleFunc("/orders", list_orders).Methods("Get")
	mux.HandleFunc("/orders", create_order).Methods("Post")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
