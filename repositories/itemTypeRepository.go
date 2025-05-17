package repositories

import (
	"errors"
	"simple-store-management/models"

	"gorm.io/gorm"
)

type ItemTypeRepository interface {
	GetAllItemTypes() (itemTypes []models.ItemType, err error)
	GetItemType(id int) (itemType models.ItemType, err error)
	CreateItemType(itemType models.ItemType) error
	UpdateItemType(itemType models.ItemType) error
	DeleteItemType(id int) error
	GetItemsOfItemType(id int) (itemType models.ItemsOfItemType, err error)
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
	if itemType.ID == 0 {
		err = errors.New("item type not found")
	}
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

func (repo *itemTypeRepository) GetItemsOfItemType(id int) (ItemOfItemType models.ItemsOfItemType, err error) {
	var itemType models.ItemTypeResponse
	err = repo.db.Table("item_type").
		Select("item_type.id AS id, item_type.type AS type").
		Joins("JOIN item ON item_type.id = item.item_type_id").
		Where("item_type.id = ?", id).
		Scan(&itemType).Error
	if err != nil {
		return
	}

	ItemOfItemType.ID = itemType.ID
	ItemOfItemType.Type = itemType.Type

	err = repo.db.Table("item").
		Select("item.id AS id, item.name AS name, item.price AS price").
		Where("item.item_type_id = ?", id).
		Scan(&ItemOfItemType.Items).Error
	return
}
