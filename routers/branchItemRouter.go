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
		api.PUT("/add", AddItemStock)
		api.PUT("/remove", RemoveItemStock)
	}
}

func CreateBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	err := branchItemSrv.CreateBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem berhasil dibuat")
}

func GetAllBranchItems(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	branchItems, err := branchItemSrv.GetAllBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branchItems berhasil diambil", branchItems)
}

func GetBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	branchItem, err := branchItemSrv.GetBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branchItem berhasil diambil", branchItem)
}

func UpdateBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	err := branchItemSrv.UpdateBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem berhasil diubah")
}

func DeleteBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	err := branchItemSrv.DeleteBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem berhasil dihapus")
}

func AddItemStock(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	err := branchItemSrv.AddItemStock(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem berhasil diubah")
}

func RemoveItemStock(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo)
	)

	err := branchItemSrv.RemoveItemStock(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem berhasil diubah")
}
