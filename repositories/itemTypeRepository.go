package repositories

import (
	"simple-store-management/models"

	"gorm.io/gorm"
)

type ItemTypeRepository interface {
	GetAllItemTypes() (itemTypes []models.ItemType, err error)
	GetItemType(id int) (itemType models.ItemType, err error)
	CreateItemType(itemType models.ItemType) error
	UpdateItemType(itemType models.ItemType) error
	DeleteItemType(id int) error
}

type itemTypeRepository struct {
	db *gorm.DB
}

func NewItemTypeRepository(database *gorm.DB) ItemTypeRepository {
	return &itemTypeRepository{
		db: database,
	}
}

func (repo *itemTypeRepository) GetAllItemTypes() (itemTypes []models.ItemType, err error) {
	err = repo.db.Find(&itemTypes).Error
	return
}

func (repo *itemTypeRepository) GetItemType(id int) (itemType models.ItemType, err error) {
	err = repo.db.Where("id = ?", id).Find(&itemType).Error
	return
}

func (repo *itemTypeRepository) CreateItemType(itemType models.ItemType) (err error) {
	err = repo.db.Create(&itemType).Error
	return
}

func (repo *itemTypeRepository) UpdateItemType(itemType models.ItemType) (err error) {
	err = repo.db.Model(&models.ItemType{}).Where("id = ?", itemType.ID).Updates(itemType).Error
	return
}

func (repo *itemTypeRepository) DeleteItemType(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&models.ItemType{}).Error
	return
}
