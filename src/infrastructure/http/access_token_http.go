package http

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/domain/access_token"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/errors"
	access_token2 "github.com/esequielvirtuoso/bookstore_oauth_api/src/services/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token2.Service
}

func NewHandler(service access_token2.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.HandleError(errors.BadRequest, errors.InvalidJSONBody)
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}