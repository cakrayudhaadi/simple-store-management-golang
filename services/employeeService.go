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
	GetTopEmployee(ctx *gin.Context) (employee models.TopEmployeeResponse, err error)
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
		err = errors.New("data employee failed to be created")
	}

	return
}

func (service *employeeService) GetAllEmployee(ctx *gin.Context) (employees []models.Employee, err error) {
	employees, err = service.employeeRepository.GetAllEmployees()
	if err != nil {
		err = errors.New("data employee failed to be loaded")
	} else if len(employees) == 0 {
		err = errors.New("data employee kosong")
	}

	return
}

func (service *employeeService) GetEmployee(ctx *gin.Context) (employee models.Employee, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	employee, err = service.employeeRepository.GetEmployee(id)

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
		err = errors.New("data employee not found")
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
		err = errors.New("data employee failed to be updated")
	}

	return
}

func (service *employeeService) DeleteEmployee(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetEmployee(ctx)
	if err != nil {
		err = errors.New("data employee not found")
		return
	}

	err = service.employeeRepository.DeleteEmployee(id)
	if err != nil {
		err = errors.New("data employee failed to be deleted")
	}

	return
}

func validateEmployeeReqAndConvertToEmployee(ctx *gin.Context) (employees models.Employee, err error) {
	var employeesRequest models.EmployeeRequest

	err = ctx.ShouldBindJSON(&employeesRequest)
	if err != nil {
		err = errors.New("parameter is not valid")
		return
	}

	err = employeesRequest.Validate()
	if err != nil {
		return
	}
	employees = employeesRequest.ConvertToEmployee()

	return
}

func (service *employeeService) GetTopEmployee(ctx *gin.Context) (topEmployee models.TopEmployeeResponse, err error) {
	var topEmployeeRequest models.TopEmployeeRequest
	topEmployeeRequest.BranchID, err = strconv.Atoi(ctx.Query("branch_id"))
	if err != nil {
		err = errors.New("parameter branch_id is required")
		return
	}
	topEmployeeRequest.Month, err = strconv.Atoi(ctx.Query("month"))
	if err != nil {
		err = errors.New("parameter month is required")
		return
	}
	topEmployeeRequest.Year, err = strconv.Atoi(ctx.Query("year"))
	if err != nil {
		err = errors.New("parameter year is required")
		return
	}

	topEmployee, err = service.employeeRepository.GetTopEmployee(topEmployeeRequest.Month, topEmployeeRequest.Year, topEmployeeRequest.BranchID)
	if err != nil {
		err = errors.New("data top employee failed to be loaded")
	}
	return
}
