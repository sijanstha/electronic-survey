package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sijanstha/electronic-voting-system/internal/adapters/repository"
)

func main() {
	profile := os.Getenv("PROFILE")
	if profile == "" || profile != "docker" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db, err := repository.NewMysqlStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewApiServer(":4000", db)
	server.Run()
}
