package services_test

import (
	"errors"
	"library-api/mocks"
	"library-api/models"
	. "library-api/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	user = models.User{
		ID:    1,
		Name:  "Dewa",
		Email: "dewa@mail.com",
		Phone: "098632137",
	}
	users     = []*models.User{&user}
	nameQuery = "dew"
)

func TestUserService_GetUsers(t *testing.T) {
	t.Run("should return list of users when no error", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		opts := &USOpts{UserRepository: mockRepo}
		s := NewUserService(opts)
		mockRepo.On("FindUsers", nameQuery).Return(users, nil)

		actualUsers, err := s.GetUsers(nameQuery)

		assert.Equal(t, users, actualUsers)
		assert.Nil(t, err)
	})

	t.Run("should return empty list and error when error get users from repository", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		opts := &USOpts{UserRepository: mockRepo}
		s := NewUserService(opts)
		dbError := errors.New("unable to connect to db")
		var expectedUsers []*models.User
		mockRepo.On("FindUsers", nameQuery).Return(expectedUsers, dbError)

		actualUsers, err := s.GetUsers(nameQuery)

		assert.Equal(t, dbError, err)
		assert.Equal(t, expectedUsers, actualUsers)
	})
}
