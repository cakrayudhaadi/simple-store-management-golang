package services

import (
	"errors"
	"simple-store-management/middlewares"
	"simple-store-management/models"
	"simple-store-management/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BranchService interface {
	CreateBranch(ctx *gin.Context) (err error)
	GetAllBranch(ctx *gin.Context) (branchs []models.Branch, err error)
	GetBranch(ctx *gin.Context) (branch models.Branch, err error)
	UpdateBranch(ctx *gin.Context) (err error)
	DeleteBranch(ctx *gin.Context) (err error)
	GetBranchWithEmployees(ctx *gin.Context) (branch models.EmployeesOfBranchResponse, err error)
	GetBranchWithItems(ctx *gin.Context) (branch models.ItemsOfBranchResponse, err error)
	GetTopBranch(ctx *gin.Context) (branch models.TopBranchResponse, err error)
}

type branchService struct {
	branchRepository repositories.BranchRepository
}

func NewBranchService(branchRepository repositories.BranchRepository) BranchService {
	return &branchService{
		branchRepository,
	}
}

func (service *branchService) CreateBranch(ctx *gin.Context) (err error) {
	var newBranch models.Branch

	newBranch, err = validateBranchReqAndConvertToBranch(ctx)
	if err != nil {
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newBranch.CreatedBy = loginName
	newBranch.CreatedAt = time.Now()

	err = service.branchRepository.CreateBranch(newBranch)
	if err != nil {
		err = errors.New("data branch gagal dibuat")
	}

	return
}

func (service *branchService) GetAllBranch(ctx *gin.Context) (branchs []models.Branch, err error) {
	branchs, err = service.branchRepository.GetAllBranchs()
	if err != nil {
		err = errors.New("data branch gagal diambil")
	} else if len(branchs) == 0 {
		err = errors.New("data branch kosong")
	}

	return
}

func (service *branchService) GetBranch(ctx *gin.Context) (branch models.Branch, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	branch, err = service.branchRepository.GetBranch(id)
	if branch.ID == 0 {
		err = errors.New("data branch tidak ada")
	} else if err != nil {
		err = errors.New("data branch gagal diambil")
	}

	return
}

func (service *branchService) UpdateBranch(ctx *gin.Context) (err error) {
	var newBranch models.Branch
	id, _ := strconv.Atoi(ctx.Param("id"))

	newBranch, err = validateBranchReqAndConvertToBranch(ctx)
	if err != nil {
		return
	}

	oldBranch, err := service.GetBranch(ctx)
	if err != nil {
		err = errors.New("data branch tidak ditemukan")
		return
	}
	newBranch.ID = id
	newBranch.CreatedBy = oldBranch.CreatedBy
	newBranch.CreatedAt = oldBranch.CreatedAt

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newBranch.UpdatedBy = loginName
	newBranch.UpdatedAt = time.Now()

	err = service.branchRepository.UpdateBranch(newBranch)
	if err != nil {
		err = errors.New("data branch gagal diubah")
	}

	return
}

func (service *branchService) DeleteBranch(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetBranch(ctx)
	if err != nil {
		err = errors.New("data branch tidak ditemukan")
		return
	}

	err = service.branchRepository.DeleteBranch(id)
	if err != nil {
		err = errors.New("data branch gagal dihapus")
	}

	return
}

func validateBranchReqAndConvertToBranch(ctx *gin.Context) (branchs models.Branch, err error) {
	var branchsRequest models.BranchRequest

	err = ctx.ShouldBindJSON(&branchsRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = branchsRequest.Validate()
	if err != nil {
		return
	}
	branchs = branchsRequest.ConvertToBranch()

	return
}

func (service *branchService) GetBranchWithEmployees(ctx *gin.Context) (branch models.EmployeesOfBranchResponse, err error) {
	return
}

func (service *branchService) GetBranchWithItems(ctx *gin.Context) (branch models.ItemsOfBranchResponse, err error) {
	return
}

func (service *branchService) GetTopBranch(ctx *gin.Context) (branch models.TopBranchResponse, err error) {
	return
}
