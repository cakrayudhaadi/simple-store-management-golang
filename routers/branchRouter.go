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
