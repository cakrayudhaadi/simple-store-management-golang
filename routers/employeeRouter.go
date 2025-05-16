package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func employeeInitiator(router *gin.Engine) {
	api := router.Group("/api/employee")
	{
		api.POST("", CreateEmployee)
		api.GET("", GetAllEmployees)
		api.GET("/:id", GetEmployee)
		api.PUT("/:id", UpdateEmployee)
		api.DELETE("/:id", DeleteEmployee)
		api.GET("/top/:branchId", GetTopEmployee)
	}
}

func CreateEmployee(ctx *gin.Context) {
	var (
		employeeRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		employeeSrv  = services.NewEmployeeService(employeeRepo)
	)

	err := employeeSrv.CreateEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data employee successfully created")
}

func GetAllEmployees(ctx *gin.Context) {
	var (
		employeeRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		employeeSrv  = services.NewEmployeeService(employeeRepo)
	)

	employees, err := employeeSrv.GetAllEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data employees successfully loaded", employees)
}

func GetEmployee(ctx *gin.Context) {
	var (
		employeeRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		employeeSrv  = services.NewEmployeeService(employeeRepo)
	)

	employee, err := employeeSrv.GetEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data employee successfully loaded", employee)
}

func UpdateEmployee(ctx *gin.Context) {
	var (
		employeeRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		employeeSrv  = services.NewEmployeeService(employeeRepo)
	)

	err := employeeSrv.UpdateEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data employee successfully updated")
}

func DeleteEmployee(ctx *gin.Context) {
	var (
		employeeRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		employeeSrv  = services.NewEmployeeService(employeeRepo)
	)

	err := employeeSrv.DeleteEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data employee successfully deleted")
}

func GetTopEmployee(ctx *gin.Context) {
	var (
		employeeRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		employeeSrv  = services.NewEmployeeService(employeeRepo)
	)

	employee, err := employeeSrv.GetTopEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data employee successfully loaded", employee)
}
