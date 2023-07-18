package main

import (
	"fmt"
	"net/http"
)

func list_products(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-type", "application/json");
  fmt.Fprint(w, `{"message":"Hello Products!"}`)
}

func list_users(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-type", "application/json");
  fmt.Fprint(w, `{"message":"Hello Users!"}`)
}

func list_orders(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-type", "application/json");
  fmt.Fprint(w, `{"message":"Hello Orders!"}`)
}

func create_order(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-type", "application/json");
  w.WriteHeader(http.StatusCreated)
}
