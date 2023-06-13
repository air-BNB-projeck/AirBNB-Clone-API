package service

import (
	"alta/air-bnb/features/users"
	"alta/air-bnb/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertData(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := users.CoreUserRequest{
		FullName: "Kurniawan",
		Email: "kurnhyalcantara@gmail.com",
		Phone: "08123456789",
		Password: "supersecret",
		Birth: "1999-03-12",
		Gender: "Male",
	}

	t.Run("Insert data success", func(t *testing.T) {
		repo.On("Insert").Return(1, nil).Once()
		userId, err := repo.Insert(insertData)
		assert.Nil(t, err)
		assert.NotNil(t, userId)
		repo.AssertExpectations(t)
	})
}