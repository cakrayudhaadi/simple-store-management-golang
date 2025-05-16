package repositories

import (
	"errors"
	"fmt"
	"simple-store-management/models"
	"time"

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
	GetBranchDetail(id int) (branch models.BranchDetailResponse, err error)
	GetTopBranch(month, year int) (branch models.TopBranchResponse, err error)
	GetBranchIDByEmployeeID(employeeID int) (branchID int, err error)
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
	if branch.ID == 0 {
		err = errors.New("branch not found")
	}
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

func (repo *branchRepository) GetBranchWithEmployees(id int) (employeeOfBranch models.EmployeesOfBranchResponse, err error) {
	var branch models.BranchResponse
	err = repo.db.Table("branch").
		Select("branch.id AS id, branch.name AS name, branch.address AS address").
		Joins("JOIN employee ON branch.id = employee.branch_id").
		Where("branch.id = ?", id).
		Scan(&branch).Error
	if err != nil {
		return
	}

	employeeOfBranch.ID = branch.ID
	employeeOfBranch.Name = branch.Name
	employeeOfBranch.Address = branch.Address

	err = repo.db.Table("employee").
		Select("employee.id AS id, employee.name AS name").
		Where("employee.branch_id = ?", id).
		Scan(&employeeOfBranch.Employees).Error
	return
}

func (repo *branchRepository) GetBranchWithItems(id int) (itemOfBranch models.ItemsOfBranchResponse, err error) {
	var branch models.BranchResponse
	err = repo.db.Table("branch").
		Select("branch.id AS id, branch.name AS name, branch.address AS address").
		Joins("JOIN employee ON branch.id = employee.branch_id").
		Where("branch.id = ?", id).
		Scan(&branch).Error
	if err != nil {
		return
	}

	itemOfBranch.ID = branch.ID
	itemOfBranch.Name = branch.Name
	itemOfBranch.Address = branch.Address

	err = repo.db.Table("item").
		Select("item.id AS id, item.name AS name, item.price AS price, branch_item.stock AS stock").
		Joins("JOIN branch_item ON item.id = branch_item.item_id").
		Where("branch_item.branch_id = ?", id).
		Scan(&itemOfBranch.Items).Error
	return
}

func (repo *branchRepository) GetBranchDetail(id int) (branchWithAll models.BranchDetailResponse, err error) {
	var branch models.BranchResponse
	err = repo.db.Table("branch").
		Select("branch.id AS id, branch.name AS name, branch.address AS address").
		Joins("JOIN employee ON branch.id = employee.branch_id").
		Where("branch.id = ?", id).
		Scan(&branch).Error
	if err != nil {
		return
	}

	branchWithAll.ID = branch.ID
	branchWithAll.Name = branch.Name
	branchWithAll.Address = branch.Address

	err = repo.db.Table("employee").
		Select("employee.id AS id, employee.name AS name").
		Where("employee.branch_id = ?", id).
		Scan(&branchWithAll.Employees).Error
	if err != nil {
		return
	}

	branchWithAll.ID = branch.ID
	branchWithAll.Name = branch.Name
	branchWithAll.Address = branch.Address

	err = repo.db.Table("item").
		Select("item.id AS id, item.name AS name, item.price AS price, branch_item.stock AS stock").
		Joins("JOIN branch_item ON item.id = branch_item.item_id").
		Where("branch_item.branch_id = ?", id).
		Scan(&branchWithAll.Items).Error
	return
}

func (repo *branchRepository) GetTopBranch(month, year int) (branch models.TopBranchResponse, err error) {
	var timestampStart string
	var timestampEnd string
	if month == 0 {
		timestampStart = fmt.Sprintf("%d-01-01 00:00:00", year)
		timestampEnd = fmt.Sprintf("%d-12-31 23:59:59", year)

		err = repo.db.Table("branch").
			Select("branch.id AS id, branch.name AS name, branch.address AS address, SUM(sales_data.amount) AS total_sales, SUM(sales_data.amount * item.price) AS total_profit").
			Joins("JOIN sales_data ON branch.id = sales_data.branch_id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.sold_date >= ? AND sales_data.sold_date <= ?", timestampStart, timestampEnd).
			Group("branch.id").
			Order("total_profit DESC").
			Limit(1).
			Scan(&branch).Error
	} else {
		firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		timestampStart = fmt.Sprintf("%d-%d-01 00:00:00", year, month)
		timestampEnd = fmt.Sprintf("%d-%d-%d 23:59:59", year, month, lastOfMonth.Day())

		err = repo.db.Table("branch").
			Select("branch.id AS id, branch.name AS name, branch.address AS address, SUM(sales_data.amount) AS total_sales, SUM(sales_data.amount * item.price) AS total_profit").
			Joins("JOIN sales_data ON branch.id = sales_data.branch_id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.sold_date >= ? AND sales_data.sold_date <= ?", timestampStart, timestampEnd).
			Group("branch.id").
			Order("total_profit DESC").
			Limit(1).
			Scan(&branch).Error
	}

	return
}

func (repo *branchRepository) GetBranchIDByEmployeeID(employeeID int) (branchID int, err error) {
	err = repo.db.Table("employee").
		Select("branch_id").
		Where("id = ?", employeeID).
		Scan(&branchID).Error
	if err != nil {
		return
	}
	if branchID == 0 {
		err = errors.New("branch not found")
	}
	return
}
