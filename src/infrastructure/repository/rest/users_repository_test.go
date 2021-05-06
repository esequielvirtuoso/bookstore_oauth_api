package rest_test

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	repoRest "github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/repository/rest"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/pkg/logger"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logger.Info("about to start rest client test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}


func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstor.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"pass"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
	})

	repository := repoRest.UsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "pass")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, errors.InvalidRestClient, err.Message)
}

func TestLoginUserInvalidErrInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstor.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"pass"}`,
		RespHTTPCode: http.StatusUnauthorized,
		RespBody: `{"message": "invalid error interface", "status": "500", "error": "internal_server_error"}`,
	})

	repository := repoRest.UsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "pass")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
}

func TestLoginUserInvalidCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstor.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"pass"}`,
		RespHTTPCode: http.StatusUnauthorized,
		RespBody: `{"message": "invalid credentials", "status": 401, "error": "unauthorized"}`,
	})

	repository := repoRest.UsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "pass")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, "invalid credentials", err.Message)
	assert.EqualValues(t, "unauthorized", err.Error)
}

func TestLoginUserInvalidJSONResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstor.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"pass"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": "1", "first_name": "Esequiel", "last_name": "Virtuoso", "email": "virtuosoesequiel@gmail.com"}`,
	})

	repository := repoRest.UsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "pass")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, errors.ErrorUnmarshalUser, err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
}

func TestLoginUserSuccess(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstor.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"pass"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": 1, "first_name": "Esequiel", "last_name": "Virtuoso", "email": "virtuosoesequiel@gmail.com"}`,
	})

	repository := repoRest.UsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "pass")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.FirstName, "Esequiel")
	assert.EqualValues(t, user.LastName, "Virtuoso")
	assert.EqualValues(t, user.Email, "virtuosoesequiel@gmail.com")
}