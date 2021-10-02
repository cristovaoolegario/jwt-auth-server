package service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	*mongo.Collection
}

func ProvideUserService(db *mongo.Database) UserService {
	return UserService{db.Collection("users")}
}

