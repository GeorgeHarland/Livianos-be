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

	// Rest test endpoint
	muxRouter.HandleFunc("/rest", func(w http.ResponseWriter, r *http.Request) {
		var err error
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode("Get request!")
		} else if r.Method == http.MethodPost {
			json.NewEncoder(w).Encode("Post request!")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err != nil {
			log.Printf("Failed to encode response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":3123", muxRouter))
}
