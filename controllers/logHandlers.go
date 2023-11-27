package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khalifan-kfan/logs-app/config"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the OCAS logs analyzer!")
}

func HandleLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for key, value := range data {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	client := config.GetClient()

	collection := client.Database("logs").Collection("fluentdLogs")

	_, err := collection.InsertOne(r.Context(), data)
	if err != nil {
		http.Error(w, "Failed to store logs", http.StatusInternalServerError)
		return
	}

	// Prepare a response
	response := map[string]string{"message": "Received values successfully and stored logs"}

	// Encode the response to JSON and send it back
	sendJSONResponse(w, http.StatusOK, response)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
