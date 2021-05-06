package rest

import (
	"encoding/json"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/domain/users"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstor.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type UsersRepository struct{}

func NewRepository() RestUsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{Email: email, Password: password}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.HandleError(errors.InternalError, errors.InvalidRestClient)
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.HandleError(errors.InternalError, errors.InvalidErrorInterface)
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.HandleError(errors.InternalError, errors.ErrorUnmarshalUser)
	}
	return &user, nil
}