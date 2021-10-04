package main

import (
	"github.com/cristovaoolegario/free-auth-server/db"
	"github.com/cristovaoolegario/free-auth-server/routes"
	"github.com/cristovaoolegario/free-auth-server/service"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func initDB() *mongo.Database {
	dataBase, err := db.NewDatabase(os.Getenv("ENV"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))
	if err != nil {
		panic(err)
	}
	return dataBase
}

func initApp(db *mongo.Database) App {
	wire.Build(ProvideApp, routes.ProvideRouter, routes.ProvideUserAPI, service.ProvideUserService)
	return App{}
}
