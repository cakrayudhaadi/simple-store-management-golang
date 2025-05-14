package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func salesDataInitiator(router *gin.Engine) {
	api := router.Group("/api/salesData")
	{
		api.POST("", CreateSalesData)
		api.GET("", GetAllSalesDatas)
		api.GET("/:id", GetSalesData)
		api.PUT("/:id", UpdateSalesData)
		api.DELETE("/:id", DeleteSalesData)
	}
}

func CreateSalesData(ctx *gin.Context) {
	var (
		salesDataRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		salesDataSrv  = services.NewSalesDataService(salesDataRepo)
	)

	err := salesDataSrv.CreateSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data salesData berhasil dibuat")
}

func GetAllSalesDatas(ctx *gin.Context) {
	var (
		salesDataRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		salesDataSrv  = services.NewSalesDataService(salesDataRepo)
	)

	salesDatas, err := salesDataSrv.GetAllSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data salesDatas berhasil diambil", salesDatas)
}

func GetSalesData(ctx *gin.Context) {
	var (
		salesDataRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		salesDataSrv  = services.NewSalesDataService(salesDataRepo)
	)

	salesData, err := salesDataSrv.GetSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data salesData berhasil diambil", salesData)
}

func UpdateSalesData(ctx *gin.Context) {
	var (
		salesDataRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		salesDataSrv  = services.NewSalesDataService(salesDataRepo)
	)

	err := salesDataSrv.UpdateSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data salesData berhasil diubah")
}

func DeleteSalesData(ctx *gin.Context) {
	var (
		salesDataRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		salesDataSrv  = services.NewSalesDataService(salesDataRepo)
	)

	err := salesDataSrv.DeleteSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data salesData berhasil dihapus")
}
