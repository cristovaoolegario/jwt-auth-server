package dto

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestInsertUser_ConvertToUser(t *testing.T) {
	insertUser := InsertUser{
		Name:     "Unit test",
		Email:    "unittest@email.com",
		Password: "123456",
	}
	user := insertUser.ConvertToUser()

	assert.Equal(t, insertUser.Name, user.Name)
	assert.Equal(t, insertUser.Email, user.Email)
	assert.Equal(t, true, user.Active)
	assert.IsType(t, primitive.ObjectID{}, user.ID)
}

func TestInsertUser_Validate_ShouldReturnError_WhenNameIsEmpty(t *testing.T) {
	insertUser := InsertUser{
		Name:     "",
		Email:    "",
		Password: "",
	}

	err := insertUser.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "Name is required.", err.Error())
}
func TestInsertUser_Validate_ShouldReturnError_WhenEmailIsEmpty(t *testing.T) {
	insertUser := InsertUser{
		Name:     "Unit test",
		Email:    "",
		Password: "",
	}

	err := insertUser.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "Email is required.", err.Error())
}
func TestInsertUser_Validate_ShouldReturnError_WhenPasswordIsEmpty(t *testing.T) {
	insertUser := InsertUser{
		Name:     "Unit test",
		Email:    "unittest@email.com",
		Password: "",
	}

	err := insertUser.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "Password is required.", err.Error())
}

func TestInsertUser_Validate_ShouldNotReturnError_WhenItsAValidObject(t *testing.T) {
	insertUser := InsertUser{
		Name:     "Unit test",
		Email:    "unittest@email.com",
		Password: "123456",
	}

	err := insertUser.Validate()
	assert.Nil(t, err)
}
