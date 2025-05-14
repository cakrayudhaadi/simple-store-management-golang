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

type BranchItemService interface {
	CreateBranchItem(ctx *gin.Context) (err error)
	GetAllBranchItem(ctx *gin.Context) (branchItems []models.BranchItem, err error)
	GetBranchItem(ctx *gin.Context) (branchItems models.BranchItem, err error)
	UpdateBranchItem(ctx *gin.Context) (err error)
	DeleteBranchItem(ctx *gin.Context) (err error)
	AddItemStock(ctx *gin.Context) (err error)
	RemoveItemStock(ctx *gin.Context) (err error)
}

type branchItemService struct {
	branchItemRepository repositories.BranchItemRepository
}

func NewBranchItemService(branchItemRepository repositories.BranchItemRepository) BranchItemService {
	return &branchItemService{
		branchItemRepository,
	}
}

func (service *branchItemService) CreateBranchItem(ctx *gin.Context) (err error) {
	var newBranchItem models.BranchItem

	newBranchItem, err = validateAddBranchItemReqAndConvertToBranchItem(ctx)
	if err != nil {
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newBranchItem.CreatedBy = loginName
	newBranchItem.CreatedAt = time.Now()

	err = service.branchItemRepository.CreateBranchItem(newBranchItem)
	if err != nil {
		err = errors.New("data branchItem gagal dibuat")
	}

	return
}

func (service *branchItemService) GetAllBranchItem(ctx *gin.Context) (branchItems []models.BranchItem, err error) {
	branchItems, err = service.branchItemRepository.GetAllBranchItems()
	if err != nil {
		err = errors.New("data branchItem gagal diambil")
	} else if len(branchItems) == 0 {
		err = errors.New("data branchItem kosong")
	}

	return
}

func (service *branchItemService) GetBranchItem(ctx *gin.Context) (branchItem models.BranchItem, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	branchItem, err = service.branchItemRepository.GetBranchItem(id)
	if branchItem.ID == 0 {
		err = errors.New("data branchItem tidak ada")
	} else if err != nil {
		err = errors.New("data branchItem gagal diambil")
	}

	return
}

func (service *branchItemService) UpdateBranchItem(ctx *gin.Context) (err error) {
	var newBranchItem models.BranchItem
	id, _ := strconv.Atoi(ctx.Param("id"))

	newBranchItem, err = validateAddBranchItemReqAndConvertToBranchItem(ctx)
	if err != nil {
		return
	}

	oldBranchItem, err := service.GetBranchItem(ctx)
	if err != nil {
		err = errors.New("data branchItem tidak ditemukan")
		return
	}
	newBranchItem.ID = id
	newBranchItem.CreatedBy = oldBranchItem.CreatedBy
	newBranchItem.CreatedAt = oldBranchItem.CreatedAt

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newBranchItem.UpdatedBy = loginName
	newBranchItem.UpdatedAt = time.Now()

	err = service.branchItemRepository.UpdateBranchItem(newBranchItem)
	if err != nil {
		err = errors.New("data branchItem gagal diubah")
	}

	return
}

func (service *branchItemService) DeleteBranchItem(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetBranchItem(ctx)
	if err != nil {
		err = errors.New("data branchItem tidak ditemukan")
		return
	}

	err = service.branchItemRepository.DeleteBranchItem(id)
	if err != nil {
		err = errors.New("data branchItem gagal dihapus")
	}

	return
}

func validateAddBranchItemReqAndConvertToBranchItem(ctx *gin.Context) (branchItems models.BranchItem, err error) {
	var branchItemsRequest models.AddBranchItemRequest

	err = ctx.ShouldBindJSON(&branchItemsRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = branchItemsRequest.Validate()
	if err != nil {
		return
	}
	branchItems = branchItemsRequest.ConvertToBranchItem()

	return
}

func validateRemoveBranchItemReqAndConvertToBranchItem(ctx *gin.Context) (branchItems models.BranchItem, err error) {
	var branchItemsRequest models.RemoveBranchItemRequest

	err = ctx.ShouldBindJSON(&branchItemsRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = branchItemsRequest.Validate()
	if err != nil {
		return
	}
	branchItems = branchItemsRequest.ConvertToBranchItem()

	return
}

func (service *branchItemService) AddItemStock(ctx *gin.Context) (err error) {
	return
}

func (service *branchItemService) RemoveItemStock(ctx *gin.Context) (err error) {
	return
}
