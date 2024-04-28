package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	sv "server/services"
)

func main() {
	// Create a new instance of your gRPC server struct
	pieFireDireServer := sv.NewPieFireDireServer()

	// Register the gRPC server implementation
	http.HandleFunc("/beef/summary", func(w http.ResponseWriter, r *http.Request) {
		// Call the gRPC server method and get the response
		resp, err := pieFireDireServer.Summary(r.Context(), nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error calling FindMeat: %v", err), http.StatusInternalServerError)
			return
		}

		// Convert the response to JSON
		jsonData, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshalling response to JSON: %v", err), http.StatusInternalServerError)
			return
		}

		// Set the response header
		w.Header().Set("Content-Type", "application/json")
		// Send the JSON response
		w.Write(jsonData)
	})

	// Start the HTTP server
	fmt.Println("HTTP server listening on port 8080...")
	fmt.Println("http://127.0.0.1:8080/beef/summary")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
