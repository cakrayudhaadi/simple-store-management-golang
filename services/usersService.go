package services

import (
	"errors"
	"simple-store-management/commons"
	"simple-store-management/middlewares"
	"simple-store-management/models"
	"simple-store-management/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(ctx *gin.Context) (result models.LoginResponse, err error)
	SignUp(ctx *gin.Context) (err error)
}

type userService struct {
	userRepository repositories.UsersRepository
}

func NewUsersService(userRepository repositories.UsersRepository) UserService {
	return &userService{
		userRepository,
	}
}

func (service *userService) Login(ctx *gin.Context) (result models.LoginResponse, err error) {
	var userReq models.LoginRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.Validate()
	if err != nil {
		return
	}

	user, err := service.userRepository.Login(userReq)
	if err != nil {
		return
	}

	if commons.IsValueEmpty(user.ID) {
		err = errors.New("account is not valid")
		return
	}

	matches := commons.CheckPassword(user.Password, userReq.Password)
	if !matches {
		err = errors.New("combination of username and password is not valid")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	middlewares.LoginRedis[jwtToken] = middlewares.UserLoginRedis{
		UserId:    0,
		Username:  user.Username,
		LoginAt:   time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 1),
	}

	result.Token = jwtToken

	return
}

func (service *userService) SignUp(ctx *gin.Context) (err error) {
	var userReq models.SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		err = errors.New("parameter is not valid")
		return
	}

	err = userReq.Validate()
	if err != nil {
		return
	}

	user, err := userReq.ConvertToModelForSignUp()
	if err != nil {
		return
	}

	err = service.userRepository.SignUp(user)
	if err != nil {
		err = errors.New("sign up failed")
		return
	}

	return
}
