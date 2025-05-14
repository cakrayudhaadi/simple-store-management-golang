package middlewares

import (
	"errors"
	"net/http"
	"simple-store-management/commons"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var LoginRedis = make(map[string]UserLoginRedis)

type UserLoginRedis struct {
	UserId    int64
	Username  string
	Role      string
	LoginAt   time.Time
	ExpiredAt time.Time
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			commons.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}

		data, ok := LoginRedis[tokenString]
		if !ok {
			commons.ResponseError(c, http.StatusBadRequest, "token invalid, please log in again")
			return
		}

		if time.Now().After(data.ExpiredAt) {
			commons.ResponseError(c, http.StatusBadRequest, "token expired, please log in again")
			return
		}

		c.Set("auth", data)

		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if commons.IsValueEmpty(authHeader) {
		return tokenString, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}

func GenerateJwtToken() (token string, err error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &jwt.RegisteredClaims{
		Issuer:    "quiz3",
		ExpiresAt: &jwt.NumericDate{Time: expirationTime},
	}

	generatedTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = generatedTokenJwt.SignedString([]byte(viper.GetString("jwt_secret_key")))
	if err != nil {
		return
	}

	return
}

func GetUsernameFromToken(c *gin.Context) (username string, err error) {
	tokenString, err := GetJwtTokenFromHeader(c)
	if err != nil {
		commons.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	data, ok := LoginRedis[tokenString]
	if !ok {
		err = errors.New("token invalid, please log in again")
		commons.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	username = data.Username
	if commons.IsValueEmpty(username) {
		err = errors.New("username empty, please log in again")
		commons.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	return
}
