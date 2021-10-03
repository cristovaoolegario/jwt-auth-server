package dto

import (
	"github.com/cristovaoolegario/free-auth-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type InsertUser struct {
	Name     string ` json:"name"`
	Email    string ` json:"email"`
	Password string ` json:"-"`
}

func (user *InsertUser) ConvertToUser() models.User {
	return models.User{
		ID:     primitive.NewObjectID(),
		Name: user.Name,
		Email:    user.Email,
		Password: getEncryptedPassword(user.Password),
		Active: true,
	}
}

func getEncryptedPassword(password string) []byte {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil{
		panic(err)
	}
	return result
}
