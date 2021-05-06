package app

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/domain/access_token"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/clients/cassandra"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/http"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	/*
	* DB
	*/
	cassandra.GetSession()

	/*
	 * Storages
	 */
	dbRepository := db.NewRepository()

	/*
	 * Services
	 */
	atService := access_token.NewService(dbRepository)

	/*
	 * Handler
	 */
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}