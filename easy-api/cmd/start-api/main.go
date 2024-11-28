package main

import (
	"log"
	"os"

	"github.com/ikiwq/ackme/easy-api/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	connStr := os.Getenv("CONN_STR")

	api := api.NewApi("localhost", "7878", connStr)

	api.Start()
}