package app

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/clients/cassandra"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/http"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/repository/db"
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/infrastructure/repository/rest"
	access_token2 "github.com/esequielvirtuoso/bookstore_oauth_api/src/services/access_token"
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
	repoRepository := rest.NewRepository()

	/*
	 * Services
	 */
	atService := access_token2.NewService(dbRepository, repoRepository)

	/*
	 * Handler
	 */
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}