package controllers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the OCAS analyzer!")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Handle the POST request here
	// For simplicity, let's just echo back the received data
	message := r.FormValue("message")
	fmt.Fprintf(w, "Received message: %s", message)
}
