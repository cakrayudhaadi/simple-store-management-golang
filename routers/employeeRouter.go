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
	}
}

func CreateEmployee(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		categorySrv  = services.NewEmployeeService(categoryRepo)
	)

	err := categorySrv.CreateEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllEmployees(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		categorySrv  = services.NewEmployeeService(categoryRepo)
	)

	books, err := categorySrv.GetAllEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetEmployee(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		categorySrv  = services.NewEmployeeService(categoryRepo)
	)

	book, err := categorySrv.GetEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateEmployee(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		categorySrv  = services.NewEmployeeService(categoryRepo)
	)

	err := categorySrv.UpdateEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteEmployee(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewEmployeeRepository(connection.DBConnections)
		categorySrv  = services.NewEmployeeService(categoryRepo)
	)

	err := categorySrv.DeleteEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
