package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func list_products(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, `{"message":"Hello Products!"}`)
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
