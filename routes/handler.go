package routes

import (
	"net/http"

	"github.com/khalifan-kfan/logs-app/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.HelloHandler)
	http.HandleFunc("/api/v1/handle_logs", controllers.HandleLogs)
}
