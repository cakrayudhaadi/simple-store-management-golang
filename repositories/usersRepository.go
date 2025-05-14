package repositories

import (
	"errors"
	"simple-store-management/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Login(user models.LoginRequest) (result models.Users, err error)
	SignUp(user models.Users) (err error)
	DeleteUsers(user models.Users) (err error)
	GetListUsers() (users []models.Users, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(database *gorm.DB) UsersRepository {
	return &userRepository{
		db: database,
	}
}

func (repo *userRepository) Login(user models.LoginRequest) (result models.Users, err error) {
	err = repo.db.Where("username = ?", user.Username).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}

	return result, nil
}

func (repo *userRepository) SignUp(user models.Users) (err error) {
	err = repo.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) DeleteUsers(user models.Users) (err error) {
	err = repo.db.Where("username = ?", user.Username).Delete(&models.Users{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) GetListUsers() (users []models.Users, err error) {
	err = repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
