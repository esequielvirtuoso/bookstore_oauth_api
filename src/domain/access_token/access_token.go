package access_token

import (
	"fmt"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	"github.com/esequielvirtuoso/bookstore_users_api/pkg/crypto_utils"
	"strings"
	"time"
)

const (
	ExpirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client credentials grant_type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type LoginPass struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		if at.Username == "" || at.Password == "" {
			return errors.HandleError(errors.Unauthorized, errors.Unauthorized)
		}
		break
	case grantTypeClientCredentials:
		if at.ClientID == "" || at.ClientSecret == "" {
			return errors.HandleError(errors.Unauthorized, errors.Unauthorized)
		}
		break
	default:
		return errors.HandleError(errors.BadRequest, errors.InvalidGrantType)
	}

	return nil
}

func (at *AccessTokenRequest) GetCredentials() (*LoginPass, *errors.RestErr) {
	var loginData LoginPass
	switch at.GrantType {
	case grantTypePassword:
		loginData.Login = at.Username
		loginData.Pass = at.Password
		break
	case grantTypeClientCredentials:
		loginData.Login = at.ClientID
		loginData.Pass = at.ClientSecret
		break
	default:
		return nil, errors.HandleError(errors.BadRequest, errors.InvalidGrantType)
	}

	return &loginData, nil
}

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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(ExpirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}