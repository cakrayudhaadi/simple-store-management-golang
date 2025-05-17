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

// Login godoc
// @Summary      Login
// @Description  login
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 request body models.LoginRequest true "request body"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Failure      500  {object}  commons.SwaggerApiResponseError
// @Router       /api/users/login [post]
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

// SignUp godoc
// @Summary      Sign Up
// @Description  sign up
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 request body models.SignUpRequest true "request body"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Failure      500  {object}  commons.SwaggerApiResponseError
// @Router       /api/users/signup [post]
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
