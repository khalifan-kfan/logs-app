package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/khalifan-kfan/logs-app/config"
	"github.com/khalifan-kfan/logs-app/routes"
)

func main() {
	routes.SetupRoutes()

	client, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server listening on port 8089...")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
