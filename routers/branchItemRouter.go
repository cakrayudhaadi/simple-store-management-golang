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
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo, branchRepo, itemRepo)
	)

	err := branchItemSrv.CreateBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem successfully created")
}

func GetAllBranchItems(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo, branchRepo, itemRepo)
	)

	branchItems, err := branchItemSrv.GetAllBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branchItems successfully loaded", branchItems)
}

func GetBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo, branchRepo, itemRepo)
	)

	branchItem, err := branchItemSrv.GetBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branchItem successfully loaded", branchItem)
}

func UpdateBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo, branchRepo, itemRepo)
	)

	err := branchItemSrv.UpdateBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem successfully updated")
}

func DeleteBranchItem(ctx *gin.Context) {
	var (
		branchItemRepo = repositories.NewBranchItemRepository(connection.DBConnections)
		branchRepo     = repositories.NewBranchRepository(connection.DBConnections)
		itemRepo       = repositories.NewItemRepository(connection.DBConnections)
		branchItemSrv  = services.NewBranchItemService(branchItemRepo, branchRepo, itemRepo)
	)

	err := branchItemSrv.DeleteBranchItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branchItem successfully deleted")
}
