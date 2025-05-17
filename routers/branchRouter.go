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

func branchInitiator(router *gin.Engine) {
	api := router.Group("/api/branch")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateBranch)
		api.GET("", GetAllBranchs)
		api.GET("/:id", GetBranch)
		api.PUT("/:id", UpdateBranch)
		api.DELETE("/:id", DeleteBranch)
		api.GET("/employees/:id", GetBranchWithEmployees)
		api.GET("/items/:id", GetBranchWithItems)
		api.GET("/detail/:id", GetBranchDetail)
		api.GET("/top", GetTopBranch)
	}
}

// Create Branch godoc
// @Summary      Create Branch
// @Description  create a new branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param 		 branchRequest body models.BranchRequest true "Branch Request"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch [post]
func CreateBranch(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	err := branchSrv.CreateBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branch successfully created")
}

// Get All Branchs godoc
// @Summary      Get All Branch
// @Description  get all branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch [get]
func GetAllBranchs(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	branchs, err := branchSrv.GetAllBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branchs successfully loaded", branchs)
}

// Get Branch godoc
// @Summary      Get Branch
// @Description  get branch by id
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/{id} [get]
func GetBranch(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	branch, err := branchSrv.GetBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch successfully loaded", branch)
}

// Get Branch Detail godoc
// @Summary      Get Branch Detail
// @Description  get branch detail
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/detail/{id} [get]
func GetBranchDetail(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	branch, err := branchSrv.GetBranchDetail(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch successfully loaded", branch)
}

// Update Branch godoc
// @Summary      Update Branch
// @Description  update a branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param 		 branchRequest body models.BranchRequest true "Branch Request"
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/{id} [put]
func UpdateBranch(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	err := branchSrv.UpdateBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branch successfully updated")
}

// Delete Branch godoc
// @Summary      Delete Branch
// @Description  delete a branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithoutData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/{id} [delete]
func DeleteBranch(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	err := branchSrv.DeleteBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branch successfully deleted")
}

// Get Branch With Employees godoc
// @Summary      Get Branch With Employees
// @Description  get branch with employees
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/employees/{id} [get]
func GetBranchWithEmployees(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	branch, err := branchSrv.GetBranchWithEmployees(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch successfully loaded", branch)
}

// Get Branch With Items godoc
// @Summary      Get Branch With Items
// @Description  get branch with Iitems
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/items/{id} [get]
func GetBranchWithItems(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	branch, err := branchSrv.GetBranchWithItems(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch successfully loaded", branch)
}

// Get Top Branch godoc
// @Summary      Get Top Branch
// @Description  get top branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        month   query      int  true  "Month"
// @Param        year   query      int  true  "Year"
// @Success      200  {object}  commons.SwaggerApiResponseSuccessWithData
// @Failure      400  {object}  commons.SwaggerApiResponseError
// @Security  	 Bearer
// @Router       /api/branch/top [get]
func GetTopBranch(ctx *gin.Context) {
	var (
		branchRepo = repositories.NewBranchRepository(connection.DBConnections)
		branchSrv  = services.NewBranchService(branchRepo)
	)

	branch, err := branchSrv.GetTopBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch successfully loaded", branch)
}
