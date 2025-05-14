package routers

import (
	"net/http"
	"simple-store-management/commons"
	"simple-store-management/databases/connection"
	"simple-store-management/repositories"
	"simple-store-management/services"

	"github.com/gin-gonic/gin"
)

func branchInitiator(router *gin.Engine) {
	api := router.Group("/api/branch")
	{
		api.POST("", CreateBranch)
		api.GET("", GetAllBranchs)
		api.GET("/:id", GetBranch)
		api.PUT("/:id", UpdateBranch)
		api.DELETE("/:id", DeleteBranch)
		api.GET("/employees", GetBranchWithEmployees)
		api.GET("/items", GetBranchWithItems)
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

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branch berhasil dibuat")
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branchs berhasil diambil", branchs)
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch berhasil diambil", branch)
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

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branch berhasil diubah")
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

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data branch berhasil dihapus")
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch berhasil diambil", branch)
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch berhasil diambil", branch)
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

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data branch berhasil diambil", branch)
}
