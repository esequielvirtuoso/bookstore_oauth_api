package access_token

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/domain/access_token"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/repository/db"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/repository/rest"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken ,*errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUserRepo rest.RestUsersRepository
	dbRepo db.DBRepository
}

func NewService(repo db.DBRepository, userRepo rest.RestUsersRepository) Service {
	return &service{dbRepo: repo, restUserRepo: userRepo}
}

func (s *service) GetById(at string) (*access_token.AccessToken, *errors.RestErr) {
	at = strings.TrimSpace(at)
	if len(at) == 0 {
		return nil, errors.HandleError(errors.BadRequest, errors.InvalidATId)
	}

	accessToken, err := s.dbRepo.GetById(at)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken ,*errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	credentials, err := request.GetCredentials()
	if err != nil {
		return nil, err
	}

	// Authenticate the user against the Users API
	user, err := s.restUserRepo.LoginUser(credentials.Login, credentials.Pass)
	if err != nil {
		return nil, err
	}

	// Generate a new access token
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

//func (s *service) Create(at AccessToken) *errors.RestErr {
//	if err := at.Validate(); err != nil {
//		return err
//	}
//	return s.dbRepo.Create(at)
//}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}