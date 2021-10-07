package service

import (
	"context"
	"github.com/cristovaoolegario/free-auth-server/dto"
	"github.com/cristovaoolegario/free-auth-server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	*mongo.Collection
}

func ProvideUserService(db *mongo.Database) UserService {
	return UserService{db.Collection("users")}
}

func (service *UserService) CreateNewUser(user dto.InsertUser) (*models.User, error) {
	userToInsert := user.ConvertToUser()
	_, err := service.InsertOne(context.TODO(), &userToInsert)
	if err != nil {
		return nil, err
	}
	return &userToInsert, nil
}
