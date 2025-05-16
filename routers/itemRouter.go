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
		itemRepo     = repositories.NewItemRepository(connection.DBConnections)
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemSrv      = services.NewItemService(itemRepo, itemTypeRepo)
	)

	err := itemSrv.CreateItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data item successfully created")
}

func GetAllItems(ctx *gin.Context) {
	var (
		itemRepo     = repositories.NewItemRepository(connection.DBConnections)
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemSrv      = services.NewItemService(itemRepo, itemTypeRepo)
	)

	items, err := itemSrv.GetAllItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data items successfully loaded", items)
}

func GetItem(ctx *gin.Context) {
	var (
		itemRepo     = repositories.NewItemRepository(connection.DBConnections)
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemSrv      = services.NewItemService(itemRepo, itemTypeRepo)
	)

	item, err := itemSrv.GetItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data item successfully loaded", item)
}

func UpdateItem(ctx *gin.Context) {
	var (
		itemRepo     = repositories.NewItemRepository(connection.DBConnections)
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemSrv      = services.NewItemService(itemRepo, itemTypeRepo)
	)

	err := itemSrv.UpdateItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data item successfully updated")
}

func DeleteItem(ctx *gin.Context) {
	var (
		itemRepo     = repositories.NewItemRepository(connection.DBConnections)
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemSrv      = services.NewItemService(itemRepo, itemTypeRepo)
	)

	err := itemSrv.DeleteItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data item successfully deleted")
}
