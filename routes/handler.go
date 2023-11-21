package routes

import (
	"net/http"

	"../controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.HelloHandler)
	http.HandleFunc("/handle_logs", controllers.PostHandler)
}
