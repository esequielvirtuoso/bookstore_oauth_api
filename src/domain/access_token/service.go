package access_token

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetById(at string) (*AccessToken, *errors.RestErr) {
	at = strings.TrimSpace(at)
	if len(at) == 0 {
		return nil, errors.HandleError(errors.BadRequest, errors.InvalidATId)
	}

	accessToken, err := s.repository.GetById(at)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}
func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}