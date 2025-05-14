package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func itemInitiator(router *gin.Engine) {
	api := router.Group("/api/item")
	{
		api.POST("", CreateItem)
		api.GET("", GetAllItems)
		api.GET("/:id", GetItem)
		api.PUT("/:id", UpdateItem)
		api.DELETE("/:id", DeleteItem)
	}
}

func CreateItem(ctx *gin.Context) {
	var (
		itemRepo = repositories.NewItemRepository(connection.DBConnections)
		itemSrv  = services.NewItemService(itemRepo)
	)

	err := itemSrv.CreateItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data item berhasil dibuat")
}

func GetAllItems(ctx *gin.Context) {
	var (
		itemRepo = repositories.NewItemRepository(connection.DBConnections)
		itemSrv  = services.NewItemService(itemRepo)
	)

	items, err := itemSrv.GetAllItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data items berhasil diambil", items)
}

func GetItem(ctx *gin.Context) {
	var (
		itemRepo = repositories.NewItemRepository(connection.DBConnections)
		itemSrv  = services.NewItemService(itemRepo)
	)

	item, err := itemSrv.GetItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data item berhasil diambil", item)
}

func UpdateItem(ctx *gin.Context) {
	var (
		itemRepo = repositories.NewItemRepository(connection.DBConnections)
		itemSrv  = services.NewItemService(itemRepo)
	)

	err := itemSrv.UpdateItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data item berhasil diubah")
}

func DeleteItem(ctx *gin.Context) {
	var (
		itemRepo = repositories.NewItemRepository(connection.DBConnections)
		itemSrv  = services.NewItemService(itemRepo)
	)

	err := itemSrv.DeleteItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data item berhasil dihapus")
}
