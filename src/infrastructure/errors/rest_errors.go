package errors

import (
	"net/http"
)

const (
	BadRequest    = "bad request"
	NotFound      = "not found"
	InternalError = "clients error"
	DatabaseError = "database error"
	Unauthorized  = "unauthorized"

	InvalidErrorInterface = "invalid error interface"
	InvalidJSONBody       = "invalid json body"
	InvalidATId           = "invalid access token id"
	InvalidUserId         = "invalid users id"
	InvalidClientId       = "invalid client id"
	InvalidExpirationTime = "invalid expiration time"
	InvalidRestClient     = "invalid rest client response"
	InvalidGrantType      = "invalid grant type"
	ErrorUnmarshalUser    = "error when trying to unmarshal users response"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func HandleError(option string, message string) *RestErr {
	err := RestErr{}
	err.Message = message
	switch option {
	case BadRequest:
		err.Status = http.StatusBadRequest
		err.Error = "bad_request"
	case NotFound:
		err.Status = http.StatusNotFound
		err.Error = "not_found"
	case InternalError:
		err.Status = http.StatusInternalServerError
		err.Error = "internal_server_error"
	case Unauthorized:
		err.Status = http.StatusUnauthorized
		err.Error = "unauthorized"
	}
	return &err
}