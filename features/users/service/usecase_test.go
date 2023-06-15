package service

import (
	"alta/air-bnb/features/users"
	"alta/air-bnb/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestInsertData(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	srv := New(repo)
	t.Run("Case #1: Insert Data with Valid Payload", func(t *testing.T) {
		insertData := users.CoreUserRequest{
			FullName: "Kurniawan",
			Email: "kurnhyalcantara@gmail.com",
			Phone: "082372527221",
			Password: "Supersecret@123",
			Birth: "1999-03-12",
			Gender: "male",
		}
		repo.On("Insert", mock.Anything).Return(uint(1), nil).Once()

		res, err := srv.RegisterUser(insertData)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), res)
		repo.AssertExpectations(t)
	})

	t.Run("Case #2: Insert With Invalid Payload", func(t *testing.T) {
		// with no email and password input, keduanya required
		insertData := users.CoreUserRequest{
			FullName: "Kurniawan",
			Password: "Supersecret@123",
			Birth: "1999-03-12",
			Gender: "male",
		}
		repo.On("Insert", mock.Anything).Return(0, errors.New("error validation: ")).Once()

		res, err := srv.RegisterUser(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res)
	})
	
	t.Run("Case #3: Insert Data Password with No Match Requirement", func(t *testing.T) {
		insertData := users.CoreUserRequest{
			FullName: "Kurniawan",
			Email: "kurnhyalcantara@gmail.com",
			Phone: "082372527221",
			Password: "supersecret", //Password harus alphanumeric dan paling tidak berisi 1 simbol dan 1 huruf besar
			Birth: "1999-03-12",
			Gender: "male",
		}

		repo.On("Insert", mock.Anything).Return(0, errors.New("error input password: ")).Once()

		res, err := srv.RegisterUser(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res)
	})

	t.Run("Case #3: Insert Data Password with Error Hash Password", func(t *testing.T) {
		insertData := users.CoreUserRequest{
			FullName: "Kurniawan",
			Email: "kurnhyalcantara@gmail.com",
			Phone: "082372527221",
			Password: "supersecret", //Password harus alphanumeric dan paling tidak berisi 1 simbol dan 1 huruf besar
			Birth: "1999-03-12",
			Gender: "male",
		}

		repo.On("Insert", mock.Anything).Return(0, errors.New("error hash password: ")).Once()

		res, err := srv.RegisterUser(insertData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res)
	})
}

func TestGetUserById(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	srv := New(repo)
	t.Run("Case #1: Success Get UserById ", func(t *testing.T) {
		insertData := users.CoreUserRequest{
				FullName: "Kurniawan",
				Email: "kurnhyalcantara@gmail.com",
				Phone: "082372527221",
				Password: "supersecret", //Password harus alphanumeric dan paling tidak berisi 1 simbol dan 1 huruf besar
				Birth: "1999-03-12",
				Gender: "male",
		}
		userId, _ := srv.RegisterUser(insertData)
		expectedResult := users.Core{
			ID: userId,
			FullName: "Kurniawan",
			Email: "kurnhyalcantara@gmail.com",
			Phone: "082372527221",
			Birth: "1999-03-12",
			Gender: "male",
		}
		repo.On("Select", mock.Anything).Return(expectedResult, nil).Once()
		result, err := srv.GetUserById(userId)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
	
	t.Run("Case #2: Failed Get User by Id", func(t *testing.T) {
		repo.On("Select", mock.Anything).Return(users.Core{}, errors.New(""))

		result, err := srv.GetUserById(0)
		assert.NotNil(t, err)
		assert.Equal(t, users.Core{}, result)
	})
}

func TestEditUserById(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	srv := New(repo)
	t.Run("Case #1: Success Edit User By Id", func(t *testing.T) {
		insertData := users.CoreUserRequest{
				FullName: "Kurniawan",
				Email: "kurnhyalcantara@gmail.com",
				Phone: "082372527221",
				Password: "Supersecret@123", //Password harus alphanumeric dan paling tidak berisi 1 simbol dan 1 huruf besar
				Birth: "1999-03-12",
				Gender: "male",
		}
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		userId, _ := srv.RegisterUser(insertData)
		updateData := users.CoreUserRequest{
			FullName: "Kurniawan Revision",
			Email: "kurnhyalcantara@gmail.com",
			Phone: "082372527221",
			Password: "Supersecret@123", //Password harus alphanumeric dan paling tidak berisi 1 simbol dan 1 huruf besar
			Birth: "1999-03-12",
			Gender: "male",
		}
		repo.On("Update", mock.Anything).Return(nil).Once()
		err := srv.EditUserById(userId, updateData);
		assert.Nil(t, err)
		user, _ := srv.GetUserById(userId)
		assert.Equal(t, "Kurniaawn Revision", user.FullName)
	})
}