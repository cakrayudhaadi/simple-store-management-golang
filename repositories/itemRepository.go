package repositories

import (
	"errors"
	"simple-store-management/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems() (items []models.Item, err error)
	GetItem(id int) (item models.Item, err error)
	CreateItem(item models.Item) error
	UpdateItem(item models.Item) error
	DeleteItem(id int) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(database *gorm.DB) ItemRepository {
	return &itemRepository{
		db: database,
	}
}

func (repo *itemRepository) GetAllItems() (items []models.Item, err error) {
	err = repo.db.Find(&items).Error
	return
}

func (repo *itemRepository) GetItem(id int) (item models.Item, err error) {
	err = repo.db.Where("id = ?", id).Find(&item).Error
	if item.ID == 0 {
		err = errors.New("item not found")
	}
	return
}

func (repo *itemRepository) CreateItem(item models.Item) (err error) {
	err = repo.db.Create(&item).Error
	return
}

func (repo *itemRepository) UpdateItem(item models.Item) (err error) {
	err = repo.db.Model(&models.Item{}).Where("id = ?", item.ID).Updates(item).Error
	return
}

func (repo *itemRepository) DeleteItem(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&models.Item{}).Error
	return
}
