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

func salesDataInitiator(router *gin.Engine) {
	api := router.Group("/api/salesData")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateSalesData)
		api.GET("", GetAllSalesDatas)
		api.GET("/:id", GetSalesData)
		api.PUT("/:id", UpdateSalesData)
		api.DELETE("/:id", DeleteSalesData)
		api.GET("/branch/:branchId", GetSalesDataBranch)
		api.GET("/employee/:employeeId", GetSalesDataEmployee)
	}
}

func CreateSalesData(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	err := salesDataSrv.CreateSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data salesData successfully created")
}

func GetAllSalesDatas(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	salesDatas, err := salesDataSrv.GetAllSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data salesDatas successfully loaded", salesDatas)
}

func GetSalesData(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	salesData, err := salesDataSrv.GetSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data salesData successfully loaded", salesData)
}

func UpdateSalesData(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	err := salesDataSrv.UpdateSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data salesData successfully updated")
}

func DeleteSalesData(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	err := salesDataSrv.DeleteSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data salesData successfully deleted")
}

func GetSalesDataBranch(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	salesData, err := salesDataSrv.GetSalesDataBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data salesData successfully loaded", salesData)
}

func GetSalesDataEmployee(ctx *gin.Context) {
	var (
		salesDataRepo  = repositories.NewSalesDataRepository(connection.DBConnections)
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		employeeRepo   = repositories.NewEmployeeRepository(connection.DBConnections)
		salesDataSrv   = services.NewSalesDataService(salesDataRepo, branchItemRepo, branchRepo, itemRepo, employeeRepo)
	)

	salesData, err := salesDataSrv.GetSalesDataEmployee(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data salesData successfully loaded", salesData)
}
