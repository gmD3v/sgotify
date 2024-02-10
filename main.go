package main

import (
	"log"

	"github.com/joho/godotenv"
	"gmd3v.com/sgotify/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}
