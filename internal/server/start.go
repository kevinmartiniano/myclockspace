package server

import (
	"net/http"
)

func StartServer() {
	RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
