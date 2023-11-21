package routes

import (
	"net/http"

	"github.com/khalifan-kfan/logs-app/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.HelloHandler)
	http.HandleFunc("/handle_logs", controllers.PostHandler)
}
