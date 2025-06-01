package router

import (
	"auth-service/handler"
	"net/http"
)

func InitRouter(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", handler.Health)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/signup", handler.SignUp)
}
