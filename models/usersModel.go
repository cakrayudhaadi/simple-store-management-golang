package models

import (
	"errors"
	"regexp"
	"simple-store-management/commons"
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Users) TableName() string {
	return "users"
}

type SignUpRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"re_type_password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (s *SignUpRequest) Validate() (err error) {
	if commons.IsValueEmpty(s.Username) {
		return errors.New("username is required")
	} else if commons.IsValueEmpty(s.Password) {
		return errors.New("password is required")
	} else if commons.IsValueEmpty(s.ReTypePassword) {
		return errors.New("retype password is required")
	} else if s.ReTypePassword != s.Password {
		return errors.New("password dan retype password are not match")
	}
	re := regexp.MustCompile(`^(.{8,})$`)
	if !re.MatchString(s.Password) {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func (l *LoginRequest) Validate() (err error) {
	if commons.IsValueEmpty(l.Username) {
		return errors.New("username is required")
	} else if commons.IsValueEmpty(l.Password) {
		return errors.New("password is required")
	}
	return
}

func (s *SignUpRequest) ConvertToModelForSignUp() (user Users, err error) {
	hashedPassword, err := commons.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password failed")
		return
	}
	return Users{
		Username: s.Username,
		Password: hashedPassword,
	}, nil
}
