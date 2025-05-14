package repositories

import (
	"simple-store-management/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAllEmployees() (employees []models.Employee, err error)
	GetEmployee(id int) (employee models.Employee, err error)
	CreateEmployee(employee models.Employee) error
	UpdateEmployee(employee models.Employee) error
	DeleteEmployee(id int) error
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
