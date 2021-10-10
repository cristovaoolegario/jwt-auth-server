package service

import (
	"context"
	"github.com/cristovaoolegario/free-auth-server/dto"
	"github.com/cristovaoolegario/free-auth-server/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (service *UserService) GetUserByEmail(email string) (*models.User, error){
	user := models.User{}
	if err := service.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil{
	 	return nil, err
	}
	return &user, nil
}
