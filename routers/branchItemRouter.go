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

func branchItemInitiator(router *gin.Engine) {
	api := router.Group("/api/branchItem")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateBranchItem)
		api.GET("", GetAllBranchItems)
		api.GET("/:id", GetBranchItem)
		api.PUT("/:id", UpdateBranchItem)
		api.DELETE("/:id", DeleteBranchItem)
	}
}

// Create Branch Item godoc
// @Summary      Create Branch Item
// @Description  create a new branch item
// @Tags         branchItem
// @Accept       json
// @Produce      json
// @Param 		 branchItemRequest body models.BranchItemRequest true "Branch Item Request"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branchItem [post]
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

// Get All Branch Items godoc
// @Summary      Get All Branch Items
// @Description  get all branch items
// @Tags         branchItem
// @Accept       json
// @Produce      json
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branchItem [get]
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

// Get Branch Item godoc
// @Summary      Get Branch Item
// @Description  get branch item by id
// @Tags         branchItem
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branchItem/{id} [get]
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

// Update Branch Item godoc
// @Summary      Update Branch Item
// @Description  update a branch item
// @Tags         branchItem
// @Accept       json
// @Produce      json
// @Param 		 branchItemUpdateRequest body models.BranchItemUpdateRequest true "Branch Item Update Request"
// @Param        id   path      int  true  "Branch Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branchItem/{id} [put]
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

// Delete Branch Item godoc
// @Summary      Delete Branch Item
// @Description  delete a branch item
// @Tags         branchItem
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch Item ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branchItem/{id} [delete]
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
