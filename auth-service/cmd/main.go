package main

import (
	"auth-service/router"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router.InitRouter(mux)
}
