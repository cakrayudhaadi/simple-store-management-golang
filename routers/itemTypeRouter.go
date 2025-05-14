package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func itemTypeInitiator(router *gin.Engine) {
	api := router.Group("/api/itemType")
	{
		api.POST("", CreateItemType)
		api.GET("", GetAllItemTypes)
		api.GET("/:id", GetItemType)
		api.PUT("/:id", UpdateItemType)
		api.DELETE("/:id", DeleteItemType)
	}
}

func CreateItemType(ctx *gin.Context) {
	var (
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemTypeSrv  = services.NewItemTypeService(itemTypeRepo)
	)

	err := itemTypeSrv.CreateItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data itemType berhasil dibuat")
}

func GetAllItemTypes(ctx *gin.Context) {
	var (
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemTypeSrv  = services.NewItemTypeService(itemTypeRepo)
	)

	itemTypes, err := itemTypeSrv.GetAllItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data itemTypes berhasil diambil", itemTypes)
}

func GetItemType(ctx *gin.Context) {
	var (
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemTypeSrv  = services.NewItemTypeService(itemTypeRepo)
	)

	itemType, err := itemTypeSrv.GetItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data itemType berhasil diambil", itemType)
}

func UpdateItemType(ctx *gin.Context) {
	var (
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemTypeSrv  = services.NewItemTypeService(itemTypeRepo)
	)

	err := itemTypeSrv.UpdateItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data itemType berhasil diubah")
}

func DeleteItemType(ctx *gin.Context) {
	var (
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemTypeSrv  = services.NewItemTypeService(itemTypeRepo)
	)

	err := itemTypeSrv.DeleteItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data itemType berhasil dihapus")
}
