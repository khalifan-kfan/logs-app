package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/khalifan-kfan/logs-app/routes"
)

func main() {
	routes.SetupRoutes()

	fmt.Println("Server listening on port 8089...")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
