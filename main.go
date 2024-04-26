package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the "hello" endpoint
	router.HandleFunc("/hello", HelloHandler).Methods("GET")

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server is running on it's all working yay ... %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

// HelloHandler handles requests to the "/hello" endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Create a HelloResponse
	response := HelloResponse{
		Message: "Hello, World!",
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the content type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
