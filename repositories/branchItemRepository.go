package repositories

import (
	"simple-store-management/models"

	"gorm.io/gorm"
)

type BranchItemRepository interface {
	GetAllBranchItems() (branchItems []models.BranchItem, err error)
	GetBranchItem(id int) (branchItem models.BranchItem, err error)
	CreateBranchItem(branchItem models.BranchItem) error
	UpdateBranchItem(branchItem models.BranchItem) error
	DeleteBranchItem(id int) error
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
