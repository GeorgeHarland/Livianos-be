package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	// larger = more memory, less read/writes (improves throughput)
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleRestRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Rest endpoint reached")
	var err error
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode("Thanks for the get request")
	} else if r.Method == http.MethodPost {
		json.NewEncoder(w).Encode("Thanks for the post request")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func handleWsRequest(w http.ResponseWriter, r *http.Request) {
	// allowing any connection for now
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client successfully connected!")
	//json.NewEncoder(w).Encode("Websocket request")
}

func setupRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/rest", handleRestRequest)
	muxRouter.HandleFunc("/websocket", handleWsRequest)
}

func main() {
	log.Println("Server start")
	muxRouter := mux.NewRouter()
	setupRoutes(muxRouter)
	log.Fatal(http.ListenAndServe(":3123", muxRouter))
}
