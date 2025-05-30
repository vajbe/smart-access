package router

import (
	"auth-service/handler"
	"net/http"
)

func InitRouter(mux *http.ServeMux) {
	mux.HandleFunc("/health", handler.Health)
}
