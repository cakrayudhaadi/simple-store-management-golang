package repositories

import (
	"errors"
	"simple-store-management/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAllEmployees() (employees []models.Employee, err error)
	GetEmployee(id int) (employee models.Employee, err error)
	CreateEmployee(employee models.Employee) error
	UpdateEmployee(employee models.Employee) error
	DeleteEmployee(id int) error
	GetTopEmployee(month, year, idBranch int) (employee models.TopEmployeeResponse, err error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(database *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: database,
	}
}

func (repo *employeeRepository) GetAllEmployees() (employees []models.Employee, err error) {
	err = repo.db.Find(&employees).Error
	return
}

func (repo *employeeRepository) GetEmployee(id int) (employee models.Employee, err error) {
	err = repo.db.Where("id = ?", id).Find(&employee).Error
	if employee.ID == 0 {
		err = errors.New("employee not found")
	}
	return
}

func (repo *employeeRepository) CreateEmployee(employee models.Employee) (err error) {
	err = repo.db.Create(&employee).Error
	return
}

func (repo *employeeRepository) UpdateEmployee(employee models.Employee) (err error) {
	err = repo.db.Model(&models.Employee{}).Where("id = ?", employee.ID).Updates(employee).Error
	return
}

func (repo *employeeRepository) DeleteEmployee(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&models.Employee{}).Error
	return
}

func (repo *employeeRepository) GetTopEmployee(month, year, branchID int) (employee models.TopEmployeeResponse, err error) {
	if month == 0 {
		err = repo.db.Table("employee").
			Select("employee.id, employee.name, branch.id, branch.name, SUM(sales_data.amount) AS total_sales, SUM(sales_data.amount * item.price) AS total_profit").
			Joins("JOIN branch ON employee.branch_id = branch.id").
			Joins("JOIN sales_data ON employee.id = sales_data.employee_id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("YEAR(sales_data.sold_date) = ? AND employee.branch_id = ?", month, year, branchID).
			Group("employee.id").
			Order("total_profit DESC").
			Limit(1).
			Scan(&employee).Error
	} else {
		err = repo.db.Table("employee").
			Select("employee.id, employee.name, branch.id, branch.name, SUM(sales_data.amount) AS total_sales, SUM(sales_data.amount * item.price) AS total_profit").
			Joins("JOIN branch ON employee.branch_id = branch.id").
			Joins("JOIN sales_data ON employee.id = sales_data.employee_id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("MONTH(sales_data.sold_date) = ? AND YEAR(sales_data.sold_date) = ? AND employee.branch_id = ?", month, year, branchID).
			Group("employee.id").
			Order("total_profit DESC").
			Limit(1).
			Scan(&employee).Error
	}

	return
}
