package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/musarafik/newsy/api"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", api.HomeHandler).Methods("GET")

	port := 8080
	fmt.Printf("Server listening on :%d... \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
