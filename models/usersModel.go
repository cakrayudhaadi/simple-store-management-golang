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
		return errors.New("username harus diisi")
	} else if commons.IsValueEmpty(s.Password) {
		return errors.New("password harus diisi")
	} else if commons.IsValueEmpty(s.ReTypePassword) {
		return errors.New("retype password harus diisi")
	} else if s.ReTypePassword != s.Password {
		return errors.New("password dan retype password tidak sama")
	}
	re := regexp.MustCompile(`^(.{8,})$`)
	if !re.MatchString(s.Password) {
		return errors.New("password harus mengandung minimal 8 karakter")
	}
	return nil
}

func (l *LoginRequest) Validate() (err error) {
	if commons.IsValueEmpty(l.Username) {
		return errors.New("username harus diisi")
	} else if commons.IsValueEmpty(l.Password) {
		return errors.New("password harus diisi")
	}
	return
}

func (s *SignUpRequest) ConvertToModelForSignUp() (user Users, err error) {
	hashedPassword, err := commons.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password gagal")
		return
	}
	return Users{
		Username: s.Username,
		Password: hashedPassword,
	}, nil
}
