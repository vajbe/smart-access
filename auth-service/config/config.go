package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Print("\nPort variable is not set")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Print("\nError occurred while converting port env variable - Defaulting to 8080")
		portInt = 8080
	}

	return &Config{
		Port: portInt,
	}
}
