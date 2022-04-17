package main

import (
	"log"

	"github.com/CesarDelgadoM/register-API/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load file .env")
	}
	server := server.New()
	server.Initialize()
	server.Run("8080")
}
