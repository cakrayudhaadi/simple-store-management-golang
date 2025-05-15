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
	GetBranchWithEmployees(id int) (branch models.EmployeesOfBranchResponse, err error)
	GetBranchWithItems(id int) (branch models.ItemsOfBranchResponse, err error)
	GetTopBranch(month, year int) (branch models.TopBranchResponse, err error)
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

func (repo *branchRepository) GetBranchWithEmployees(id int) (branch models.EmployeesOfBranchResponse, err error) {
	err = repo.db.Table("branch").
		Select("branch.id, branch.name, branch.address").
		Where("branch.id = ?", id).
		Scan(&branch).Error
	if err != nil {
		return
	}

	err = repo.db.Table("employee").
		Select("employee.id, employee.name").
		Where("employee.branch_id = ?", id).
		Scan(&branch.Employees).Error
	return
}

func (repo *branchRepository) GetBranchWithItems(id int) (branch models.ItemsOfBranchResponse, err error) {
	err = repo.db.Table("branch").
		Select("branch.id, branch.name, branch.address").
		Where("branch.id = ?", id).
		Scan(&branch).Error
	if err != nil {
		return
	}

	err = repo.db.Table("item").
		Select("item.id, item.name, item.price, branch_item.stock").
		Joins("JOIN branch_item ON item.id = branch_item.item_id").
		Where("branch_item.branch_id = ?", id).
		Scan(&branch.Items).Error
	return
}

func (repo *branchRepository) GetTopBranch(month, year int) (branch models.TopBranchResponse, err error) {
	if month == 0 {
		err = repo.db.Table("branch").
			Select("branch.id, branch.name, branch.address, SUM(sales_data.amount) AS total_sales, SUM(sales_data.amount * item.price) AS total_profit").
			Joins("JOIN sales_data ON branch.id = sales_data.branch_id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("AND YEAR(sales_data.sold_date) = ?", month, year).
			Group("branch.id").
			Order("total_profit DESC").
			Limit(1).
			Scan(&branch).Error
	} else {
		err = repo.db.Table("branch").
			Select("branch.id, branch.name, branch.address, SUM(sales_data.amount) AS total_sales, SUM(sales_data.amount * item.price) AS total_profit").
			Joins("JOIN sales_data ON branch.id = sales_data.branch_id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("MONTH(sales_data.sold_date) = ? AND YEAR(sales_data.sold_date) = ?", month, year).
			Group("branch.id").
			Order("total_profit DESC").
			Limit(1).
			Scan(&branch).Error
	}

	return
}
