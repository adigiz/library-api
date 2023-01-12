package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"git.garena.com/sea-labs-id/trainers/library-api/httperror"
	"git.garena.com/sea-labs-id/trainers/library-api/mocks"
	"git.garena.com/sea-labs-id/trainers/library-api/models"
	"git.garena.com/sea-labs-id/trainers/library-api/server"
	"git.garena.com/sea-labs-id/trainers/library-api/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
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

func TestHandler_GetUsers(t *testing.T) {
	t.Run("should return code 200 with list of users when name query match", func(t *testing.T) {
		mockUService := new(mocks.UserService)
		mockUService.On("GetUsers", nameQuery).Return(users, nil)
		opts := &server.RouterOpts{
			UserService: mockUService,
		}
		expectedRes, _ := json.Marshal(users)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/users?name=%v", nameQuery), nil)
		_, rec := testutils.ServeReq(opts, req)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedRes), rec.Body.String())
	})

	t.Run("should return error res when service return error", func(t *testing.T) {
		mockUService := new(mocks.UserService)
		err := errors.New("db error") // by default will trigger error code 500
		mockUService.On("GetUsers", nameQuery).Return(users, err)
		opts := &server.RouterOpts{
			UserService: mockUService,
		}
		expectedErrBody, _ := json.Marshal(httperror.AppError{
			StatusCode: 500,
			Code:       "INTERNAL_SERVER_ERROR",
			Message:    "db error",
		})

		req, _ := http.NewRequest("GET", fmt.Sprintf("/users?name=%v", nameQuery), nil)
		_, rec := testutils.ServeReq(opts, req)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, string(expectedErrBody), rec.Body.String())
	})
}
