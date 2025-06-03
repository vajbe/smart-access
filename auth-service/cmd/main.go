package main

import (
	"auth-service/config"
	"auth-service/db"
	"auth-service/router"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router.InitRouter(mux)
	config.LoadEnv()
	db.InitDB()
	db.RunMigrations()
	port := config.GetEnv("AUTH_PORT", ":8081")
	log.Println("🚀 Auth service running on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
