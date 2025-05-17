package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/middlewares"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func employeeInitiator(router *gin.Engine) {
	api := router.Group("/api/employee")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateEmployee)
		api.GET("", GetAllEmployees)
		api.GET("/:id", GetEmployee)
		api.PUT("/:id", UpdateEmployee)
		api.DELETE("/:id", DeleteEmployee)
		api.GET("/top/:branchId", GetTopEmployee)
	}
}

// Create Employee godoc
// @Summary      Create Employee
// @Description  create employee
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param 		 employeeRequest body models.EmployeeRequest true "Employee Request"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/employee [post]
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

// Get All Employees godoc
// @Summary      Get All Employees
// @Description  get all employees
// @Tags         employee
// @Accept       json
// @Produce      json
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/employee [get]
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

// Get Employee godoc
// @Summary      Get Employee
// @Description  get employee by id
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/employee/{id} [get]
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

// Update Employee godoc
// @Summary      Update Employee
// @Description  update employee
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param 		 employeeRequest body models.EmployeeRequest true "Employee Request"
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/employee/{id} [put]
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

// Delete Employee Item godoc
// @Summary      Delete Employee
// @Description  delete employee
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/employee/{id} [delete]
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

// Get Top Employee godoc
// @Summary      Get Top Employee
// @Description  get top employee
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param        branchId   path      int  true  "Branch ID"
// @Param        month   query      int  true  "Month"
// @Param        year   query      int  true  "Year"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/employee/top/{branchId} [get]
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
