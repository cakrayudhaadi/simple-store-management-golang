package services

import (
	"errors"
	"simple-store-management/middlewares"
	"simple-store-management/models"
	"simple-store-management/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EmployeeService interface {
	CreateEmployee(ctx *gin.Context) (err error)
	GetAllEmployee(ctx *gin.Context) (employees []models.Employee, err error)
	GetEmployee(ctx *gin.Context) (employee models.Employee, err error)
	UpdateEmployee(ctx *gin.Context) (err error)
	DeleteEmployee(ctx *gin.Context) (err error)
}

type employeeService struct {
	employeeRepository repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepository repositories.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepository,
	}
}

func (service *employeeService) CreateEmployee(ctx *gin.Context) (err error) {
	var newEmployee models.Employee

	newEmployee, err = validateEmployeeReqAndConvertToEmployee(ctx)
	if err != nil {
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newEmployee.CreatedBy = loginName
	newEmployee.CreatedAt = time.Now()

	err = service.employeeRepository.CreateEmployee(newEmployee)
	if err != nil {
		err = errors.New("data employee gagal dibuat")
	}

	return
}

func (service *employeeService) GetAllEmployee(ctx *gin.Context) (employees []models.Employee, err error) {
	employees, err = service.employeeRepository.GetAllEmployees()
	if err != nil {
		err = errors.New("data employee gagal diambil")
	} else if len(employees) == 0 {
		err = errors.New("data employee kosong")
	}

	return
}

func (service *employeeService) GetEmployee(ctx *gin.Context) (employee models.Employee, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	employee, err = service.employeeRepository.GetEmployee(id)
	if employee.ID == 0 {
		err = errors.New("data employee tidak ada")
	} else if err != nil {
		err = errors.New("data employee gagal diambil")
	}

	return
}

func (service *employeeService) UpdateEmployee(ctx *gin.Context) (err error) {
	var newEmployee models.Employee
	id, _ := strconv.Atoi(ctx.Param("id"))

	newEmployee, err = validateEmployeeReqAndConvertToEmployee(ctx)
	if err != nil {
		return
	}

	oldEmployee, err := service.GetEmployee(ctx)
	if err != nil {
		err = errors.New("data employee tidak ditemukan")
		return
	}
	newEmployee.ID = id
	newEmployee.CreatedBy = oldEmployee.CreatedBy
	newEmployee.CreatedAt = oldEmployee.CreatedAt

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newEmployee.UpdatedBy = loginName
	newEmployee.UpdatedAt = time.Now()

	err = service.employeeRepository.UpdateEmployee(newEmployee)
	if err != nil {
		err = errors.New("data employee gagal diubah")
	}

	return
}

func (service *employeeService) DeleteEmployee(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetEmployee(ctx)
	if err != nil {
		err = errors.New("data employee tidak ditemukan")
		return
	}

	err = service.employeeRepository.DeleteEmployee(id)
	if err != nil {
		err = errors.New("data employee gagal dihapus")
	}

	return
}

func validateEmployeeReqAndConvertToEmployee(ctx *gin.Context) (employees models.Employee, err error) {
	var employeesRequest models.EmployeeRequest

	err = ctx.ShouldBindJSON(&employeesRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = employeesRequest.Validate()
	if err != nil {
		return
	}
	employees = employeesRequest.ConvertToEmployee()

	return
}
