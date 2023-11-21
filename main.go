package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/khalifan-kfan/logs-app/routes"
)

func main() {
	routes.SetupRoutes()

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
