package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello, World!")
	muxRouter := mux.NewRouter()

	// Rest
	muxRouter.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode("Get request!")
		} else if r.Method == http.MethodPost {
			json.NewEncoder(w).Encode("Post request!")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//// Graphql
	//h := handler.New(&handler.config{
	//	Schema:   &schema,
	//	Pretty:   true,
	//	GraphiQL: true,
	//})
	//
	//muxRouter.Handle("/graphql", h)

	http.ListenAndServe(":3123", muxRouter)
	//log.Fatal(http.ListenAndServe(":3123", muxRouter))
}
