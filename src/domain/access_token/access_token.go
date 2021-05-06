package access_token

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	"strings"
	"time"
)

const (
	ExpirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` // application that is requesting the token so we can limit the token timeout
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.HandleError(errors.BadRequest, errors.InvalidATId)
	}
	if at.UserId == 0 {
		return errors.HandleError(errors.BadRequest, errors.InvalidUserId)
	}
	if at.ClientId == 0 {
		return errors.HandleError(errors.BadRequest, errors.InvalidClientId)
	}
	if at.Expires == 0 {
		return errors.HandleError(errors.BadRequest, errors.InvalidExpirationTime)
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(ExpirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken)IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}