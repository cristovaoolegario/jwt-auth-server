package dto

import (
	"errors"
	"github.com/cristovaoolegario/free-auth-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

type InsertUser struct {
	Name     string ` json:"name"`
	Email    string ` json:"email"`
	Password string ` json:"password"`
}

func (user *InsertUser) ConvertToUser() models.User {
	return models.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: getEncryptedPassword(user.Password),
		Active:   true,
	}
}

func (user *InsertUser) Validate() error {
	if len(user.Name) == 0 {
		return MissingFieldError("Name")
	}
	if len(user.Email) == 0 {
		return MissingFieldError("Email")
	}
	if len(user.Password) == 0 {
		return MissingFieldError("Password")
	}
	if isValidEmail(user.Email) != nil {
		return errors.New("Invalid email.")
	}
	return nil
}

func MissingFieldError(missingField string) error {
	return errors.New(missingField + " is required.")
}

func isValidEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func getEncryptedPassword(password string) []byte {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return result
}
