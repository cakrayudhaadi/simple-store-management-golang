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
	}
}

func CreateBranch(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchRepository(connection.DBConnections)
		categorySrv  = services.NewBranchService(categoryRepo)
	)

	err := categorySrv.CreateBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllBranchs(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchRepository(connection.DBConnections)
		categorySrv  = services.NewBranchService(categoryRepo)
	)

	books, err := categorySrv.GetAllBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetBranch(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchRepository(connection.DBConnections)
		categorySrv  = services.NewBranchService(categoryRepo)
	)

	book, err := categorySrv.GetBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateBranch(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchRepository(connection.DBConnections)
		categorySrv  = services.NewBranchService(categoryRepo)
	)

	err := categorySrv.UpdateBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteBranch(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBranchRepository(connection.DBConnections)
		categorySrv  = services.NewBranchService(categoryRepo)
	)

	err := categorySrv.DeleteBranch(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
