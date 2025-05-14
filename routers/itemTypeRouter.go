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
		categoryRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		categorySrv  = services.NewItemTypeService(categoryRepo)
	)

	err := categorySrv.CreateItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllItemTypes(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		categorySrv  = services.NewItemTypeService(categoryRepo)
	)

	books, err := categorySrv.GetAllItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetItemType(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		categorySrv  = services.NewItemTypeService(categoryRepo)
	)

	book, err := categorySrv.GetItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateItemType(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		categorySrv  = services.NewItemTypeService(categoryRepo)
	)

	err := categorySrv.UpdateItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteItemType(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		categorySrv  = services.NewItemTypeService(categoryRepo)
	)

	err := categorySrv.DeleteItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
