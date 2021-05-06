package db

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/domain/access_token"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/clients/cassandra"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/pkg/logger"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		logger.Error("error when trying to get access token from db", err)
		if err == gocql.ErrNotFound {
			return nil, errors.HandleError(errors.NotFound, err.Error())
		}
		return nil, errors.HandleError(errors.InternalError, err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		logger.Error("error when trying to create a new access token", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		logger.Error("error when trying to update the access token expiration time", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	return nil
}