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

// Create Sales Data godoc
// @Summary      Create Sales Data
// @Description  create a new sales data
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Param 		 salesDataRequest body models.SalesDataRequest true "Sales Data Request"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData [post]
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

// Get All Sales Datas godoc
// @Summary      Get All Sales Datas
// @Description  get all sales datas
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData [get]
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

// Get Sales Data godoc
// @Summary      Get Sales Data
// @Description  get sales data by id
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Sales Data ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData/{id} [get]
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

// Update Sales Data godoc
// @Summary      Update Sales Data
// @Description  update a sales data
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Param 		 salesDataRequest body models.SalesDataRequest true "Sales Data Request"
// @Param        id   path      int  true  "Branch Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData/{id} [put]
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

// Delete Sales Data godoc
// @Summary      Delete Sales Data
// @Description  delete a sales data
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Sales Data ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData/{id} [delete]
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

// Get Sales Data Branch godoc
// @Summary      Get Sales Data Branch
// @Description  get sales data by branch
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Param        branchId   path      int  true  "Branch ID"
// @Param        month   query      int  true  "Month"
// @Param        year   query      int  true  "Year"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData/branch/{branchId} [get]
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

// Get Sales Data Employee godoc
// @Summary      Get Sales Data Employee
// @Description  get sales data by employee
// @Tags         salesData
// @Accept       json
// @Produce      json
// @Param        employeeId   path      int  true  "Employee ID"
// @Param        month   query      int  true  "Month"
// @Param        year   query      int  true  "Year"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/salesData/employee/{employeeId} [get]
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
