package repositories

import (
	"errors"
	"simple-store-management/models"

	"gorm.io/gorm"
)

type BranchItemRepository interface {
	GetAllBranchItems() (branchItems []models.BranchItem, err error)
	GetBranchItem(id int) (branchItem models.BranchItem, err error)
	CreateBranchItem(branchItem models.BranchItem) error
	UpdateBranchItem(branchItem models.BranchItem) error
	DeleteBranchItem(id int) error
	GetBranchItemByBranchIDAndItemID(branchId int, itemId int) (branchItem models.BranchItem, err error)
}

type branchItemRepository struct {
	db *gorm.DB
}

func NewBranchItemRepository(database *gorm.DB) BranchItemRepository {
	return &branchItemRepository{
		db: database,
	}
}

func (repo *branchItemRepository) GetAllBranchItems() (branchItems []models.BranchItem, err error) {
	err = repo.db.Find(&branchItems).Error
	return
}

func (repo *branchItemRepository) GetBranchItem(id int) (branchItem models.BranchItem, err error) {
	err = repo.db.Where("id = ?", id).Find(&branchItem).Error
	if branchItem.ID == 0 {
		err = errors.New("branch item not found")
	}
	return
}

func (repo *branchItemRepository) CreateBranchItem(branchItem models.BranchItem) (err error) {
	err = repo.db.Create(&branchItem).Error
	return
}

func (repo *branchItemRepository) UpdateBranchItem(branchItem models.BranchItem) (err error) {
	err = repo.db.Model(&models.BranchItem{}).Where("id = ?", branchItem.ID).Updates(branchItem).Error
	return
}

func (repo *branchItemRepository) DeleteBranchItem(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&models.BranchItem{}).Error
	return
}

func (repo *branchItemRepository) GetBranchItemByBranchIDAndItemID(branchId int, itemId int) (branchItem models.BranchItem, err error) {
	err = repo.db.Where("branch_id = ? AND item_id = ?", branchId, itemId).Find(&branchItem).Error
	if err != nil {
		return
	}
	return
}
