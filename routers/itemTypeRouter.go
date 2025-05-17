package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/middlewares"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func itemTypeInitiator(router *gin.Engine) {
	api := router.Group("/api/itemType")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateItemType)
		api.GET("", GetAllItemTypes)
		api.GET("/:id", GetItemType)
		api.PUT("/:id", UpdateItemType)
		api.DELETE("/:id", DeleteItemType)
		api.GET("/items/:id", GetItemsOfItemType)
	}
}

// Create Item Type godoc
// @Summary      Create Item Type
// @Description  create a new item type
// @Tags         itemType
// @Accept       json
// @Produce      json
// @Param 		 itemTypeRequest body models.ItemTypeRequest true "Item Type Request"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/itemType [post]
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

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data itemType successfully created")
}

// Get All Item Types godoc
// @Summary      Get All Item Types
// @Description  get all items types
// @Tags         itemType
// @Accept       json
// @Produce      json
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/itemType [get]
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data itemTypes successfully loaded", itemTypes)
}

// Get Item Type godoc
// @Summary      Get Item Type
// @Description  get item type by id
// @Tags         itemType
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item Type ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/itemType/{id} [get]
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data itemType successfully loaded", itemType)
}

// Update Item Type godoc
// @Summary      Update Item Type
// @Description  update item type by id
// @Tags         itemType
// @Accept       json
// @Produce      json
// @Param 		 itemTypeRequest body models.ItemTypeRequest true "Item Type Request"
// @Param        id   path      int  true  "Branch Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/itemType/{id} [put]
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

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data itemType successfully updated")
}

// Delete Item Type godoc
// @Summary      Delete Item Type
// @Description  delete item type by id
// @Tags         itemType
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item Type ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/itemType/{id} [delete]
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

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data itemType successfully deleted")
}

// Get Items Of Item Type godoc
// @Summary      Get Items Of Item Type
// @Description  get items of item type
// @Tags         itemType
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item Type ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/itemType/items/{id} [get]
func GetItemsOfItemType(ctx *gin.Context) {
	var (
		itemTypeRepo = repositories.NewItemTypeRepository(connection.DBConnections)
		itemTypeSrv  = services.NewItemTypeService(itemTypeRepo)
	)

	itemTypes, err := itemTypeSrv.GetItemsOfItemType(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data items successfully loaded", itemTypes)
}
