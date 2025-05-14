package repositories

import (
	"simple-store-management/models"

	"gorm.io/gorm"
)

type BranchRepository interface {
	GetAllBranchs() (branchs []models.Branch, err error)
	GetBranch(id int) (branch models.Branch, err error)
	CreateBranch(branch models.Branch) error
	UpdateBranch(branch models.Branch) error
	DeleteBranch(id int) error
}

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(database *gorm.DB) BranchRepository {
	return &branchRepository{
		db: database,
	}
}

func (repo *branchRepository) GetAllBranchs() (branchs []models.Branch, err error) {
	err = repo.db.Find(&branchs).Error
	return
}

func (repo *branchRepository) GetBranch(id int) (branch models.Branch, err error) {
	err = repo.db.Where("id = ?", id).Find(&branch).Error
	return
}

func (repo *branchRepository) CreateBranch(branch models.Branch) (err error) {
	err = repo.db.Create(&branch).Error
	return
}

func (repo *branchRepository) UpdateBranch(branch models.Branch) (err error) {
	err = repo.db.Model(&models.Branch{}).Where("id = ?", branch.ID).Updates(branch).Error
	return
}

func (repo *branchRepository) DeleteBranch(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&models.Branch{}).Error
	return
}
