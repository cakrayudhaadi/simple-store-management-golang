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

func itemInitiator(router *gin.Engine) {
	api := router.Group("/api/item")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateItem)
		api.GET("", GetAllItems)
		api.GET("/:id", GetItem)
		api.PUT("/:id", UpdateItem)
		api.DELETE("/:id", DeleteItem)
	}
}

// Create Item godoc
// @Summary      Create Item
// @Description  create a new item
// @Tags         item
// @Accept       json
// @Produce      json
// @Param 		 itemRequest body models.ItemRequest true "Item Request"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/item [post]
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

// Get All Items godoc
// @Summary      Get All Items
// @Description  get all items
// @Tags         item
// @Accept       json
// @Produce      json
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/item [get]
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

// Get Item godoc
// @Summary      Get Item
// @Description  get item by id
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/item/{id} [get]
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

// Update Item godoc
// @Summary      Update Item
// @Description  update item by id
// @Tags         item
// @Accept       json
// @Produce      json
// @Param 		 itemRequest body models.ItemRequest true "Item Request"
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/item/{id} [put]
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

// Delete Item godoc
// @Summary      Delete Item
// @Description  delete item by id
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/item/{id} [delete]
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
