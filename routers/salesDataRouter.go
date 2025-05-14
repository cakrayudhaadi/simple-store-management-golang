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
		categoryRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		categorySrv  = services.NewSalesDataService(categoryRepo)
	)

	err := categorySrv.CreateSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllSalesDatas(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		categorySrv  = services.NewSalesDataService(categoryRepo)
	)

	books, err := categorySrv.GetAllSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetSalesData(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		categorySrv  = services.NewSalesDataService(categoryRepo)
	)

	book, err := categorySrv.GetSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateSalesData(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		categorySrv  = services.NewSalesDataService(categoryRepo)
	)

	err := categorySrv.UpdateSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteSalesData(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewSalesDataRepository(connection.DBConnections)
		categorySrv  = services.NewSalesDataService(categoryRepo)
	)

	err := categorySrv.DeleteSalesData(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
