// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/cristovaoolegario/free-auth-server/db"
	"github.com/cristovaoolegario/free-auth-server/routes"
	"github.com/cristovaoolegario/free-auth-server/service"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// Injectors from wire.go:

func InitApp(db *mongo.Database) App {
	userService := service.ProvideUserService(db)
	userAPI := routes.ProvideUserAPI(userService)
	router := routes.ProvideRouter(userAPI)
	app := ProvideApp(router)
	return app
}

// wire.go:

func InitDB() *mongo.Database {
	dataBase, err := db.NewDatabase(os.Getenv("ENV"), os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))
	if err != nil {
		panic(err)
	}
	return dataBase
}
