package access_token_test

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/domain/access_token"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.Equal(t, 24, access_token.ExpirationTime)
}

func TestGetNewAccessToken(t *testing.T) {
	at := access_token.GetNewAccessToken()

	assert.NotEqual(t, true, at.IsExpired()) // brand new access token should not be expired
	assert.Empty(t, at.AccessToken)          // new access token should not have defined access token id
	assert.Zero(t, at.UserId)                // New access token should not have an associated users id
 }

 func TestAccessToken_IsExpired(t *testing.T) {
	 at := access_token.AccessToken{}
	assert.Equal(t, true, at.IsExpired()) // empty access token should be expired by default

	 at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	 assert.Equal(t, false, at.IsExpired()) // access token expiring three hours from now should not be expired
 }