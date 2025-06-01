package main

import (
	"auth-service/config"
	"auth-service/router"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router.InitRouter(mux)
	c := config.LoadConfig()
	fmt.Printf("\nStarting server on: %d", c.Port)
	portStr := fmt.Sprintf(":%d", c.Port)
	http.ListenAndServe(portStr, mux)
}
