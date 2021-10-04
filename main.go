package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := InitDB()
	app := InitApp(db)

	app.Run(os.Getenv("PORT"),
		os.Getenv("ENV"))
}
