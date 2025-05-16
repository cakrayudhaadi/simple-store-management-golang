package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func userInitiator(router *gin.Engine) {
	api := router.Group("/api/users")
	{
		api.POST("/login", Login)
		api.POST("/signup", SignUp)
	}
}

func Login(ctx *gin.Context) {
	var (
		userRepo = repositories.NewUsersRepository(connection.DBConnections)
		userSrv  = services.NewUsersService(userRepo)
	)

	token, err := userSrv.Login(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "successfully logged in", token)
}

func SignUp(ctx *gin.Context) {
	var (
		userRepo = repositories.NewUsersRepository(connection.DBConnections)
		userSrv  = services.NewUsersService(userRepo)
	)

	err := userSrv.SignUp(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "user account successfully created")
}
