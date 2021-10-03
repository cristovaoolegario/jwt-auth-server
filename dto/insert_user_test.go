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
