package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func branchItemInitiator(router *gin.Engine) {
	api := router.Group("/api/branchItem")
	{
		api.POST("", CreateBranchItem)
		api.GET("", GetAllBranchItems)
		api.GET("/:id", GetBranchItem)
		api.PUT("/:id", UpdateBranchItem)
		api.DELETE("/:id", DeleteBranchItem)
	}
}

func CreateBranchItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		categorySrv  = services.NewBranchItemService(categoryRepo)
	)

	err := categorySrv.CreateBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllBranchItems(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		categorySrv  = services.NewBranchItemService(categoryRepo)
	)

	books, err := categorySrv.GetAllBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetBranchItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		categorySrv  = services.NewBranchItemService(categoryRepo)
	)

	book, err := categorySrv.GetBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateBranchItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		categorySrv  = services.NewBranchItemService(categoryRepo)
	)

	err := categorySrv.UpdateBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteBranchItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		categorySrv  = services.NewBranchItemService(categoryRepo)
	)

	err := categorySrv.DeleteBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
