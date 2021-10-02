package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := InitDB()
	app := InitApp(db)

	app.Run(os.Getenv("PORT"),
		os.Getenv("ENV"))
}
