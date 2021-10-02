package main

import (
	"fmt"
	"github.com/cristovaoolegario/free-auth-server/routes"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type App struct {
	*routes.Router
}

func ProvideApp(router *routes.Router) App {
	return App{router}
}

func (app *App) Run(port, env string) {
	fmt.Println("Server running in port:", port)
	if env == "dev" {
		log.Print("CORS configured for DEV environment")
		corsWrapper := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
		})
		log.Fatal(http.ListenAndServe(port, corsWrapper.Handler(app.Router)))
	} else {
		log.Fatal(http.ListenAndServe(port, app.Router))
	}
}
