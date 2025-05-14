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
		categoryRepo = repositories.NewItemRepository(connection.DBConnections)
		categorySrv  = services.NewItemService(categoryRepo)
	)

	err := categorySrv.CreateItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllItems(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemRepository(connection.DBConnections)
		categorySrv  = services.NewItemService(categoryRepo)
	)

	books, err := categorySrv.GetAllItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemRepository(connection.DBConnections)
		categorySrv  = services.NewItemService(categoryRepo)
	)

	book, err := categorySrv.GetItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemRepository(connection.DBConnections)
		categorySrv  = services.NewItemService(categoryRepo)
	)

	err := categorySrv.UpdateItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteItem(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewItemRepository(connection.DBConnections)
		categorySrv  = services.NewItemService(categoryRepo)
	)

	err := categorySrv.DeleteItem(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
