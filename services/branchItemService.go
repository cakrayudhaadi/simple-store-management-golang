package services

import (
	"errors"
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
}

type branchItemService struct {
	branchItemRepository repositories.BranchItemRepository
	branchRepository     repositories.BranchRepository
	itemRepository       repositories.ItemRepository
}

func NewBranchItemService(
	branchItemRepository repositories.BranchItemRepository,
	branchRepository repositories.BranchRepository,
	itemRepository repositories.ItemRepository,
) BranchItemService {
	return &branchItemService{
		branchItemRepository,
		branchRepository,
		itemRepository,
	}
}

func (service *branchItemService) CreateBranchItem(ctx *gin.Context) (err error) {
	var newBranchItem models.BranchItem

	newBranchItem, err = validateBranchItemReqAndConvertToBranchItem(ctx)
	if err != nil {
		return
	}

	_, err = service.branchRepository.GetBranch(newBranchItem.BranchID)
	if err != nil {
		err = errors.New("data branch not found")
		return
	}

	_, err = service.itemRepository.GetItem(newBranchItem.ItemID)
	if err != nil {
		err = errors.New("data item not found")
		return
	}

	oldBranchItem, err := service.branchItemRepository.GetBranchItemByBranchIDAndItemID(newBranchItem.BranchID, newBranchItem.ItemID)
	if oldBranchItem.ID != 0 {
		err = errors.New("data branch item already exists")
	}

	// loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	// newBranchItem.CreatedBy = loginName
	newBranchItem.CreatedAt = time.Now()

	err = service.branchItemRepository.CreateBranchItem(newBranchItem)
	if err != nil {
		err = errors.New("data branch item failed to be created")
	}

	return
}

func (service *branchItemService) GetAllBranchItem(ctx *gin.Context) (branchItems []models.BranchItem, err error) {
	branchItems, err = service.branchItemRepository.GetAllBranchItems()
	if err != nil {
		err = errors.New("data branch item failed to be loaded")
	} else if len(branchItems) == 0 {
		err = errors.New("data branch item kosong")
	}

	return
}

func (service *branchItemService) GetBranchItem(ctx *gin.Context) (branchItem models.BranchItem, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	branchItem, err = service.branchItemRepository.GetBranchItem(id)

	return
}

func (service *branchItemService) UpdateBranchItem(ctx *gin.Context) (err error) {
	var newBranchItem models.BranchItem
	id, _ := strconv.Atoi(ctx.Param("id"))

	newBranchItem, err = validateBranchItemUpdateAndConvertToBranchItem(ctx)
	if err != nil {
		return
	}

	oldBranchItem, err := service.GetBranchItem(ctx)
	if err != nil {
		err = errors.New("data branch item not found")
		return
	}
	newBranchItem.BranchID = oldBranchItem.BranchID
	newBranchItem.ItemID = oldBranchItem.ItemID

	_, err = service.branchRepository.GetBranch(newBranchItem.BranchID)
	if err != nil {
		err = errors.New("data branch not found")
		return
	}

	_, err = service.itemRepository.GetItem(newBranchItem.ItemID)
	if err != nil {
		err = errors.New("data item not found")
		return
	}

	newBranchItem.ID = id
	newBranchItem.CreatedBy = oldBranchItem.CreatedBy
	newBranchItem.CreatedAt = oldBranchItem.CreatedAt

	// loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	// newBranchItem.UpdatedBy = loginName
	newBranchItem.UpdatedAt = time.Now()

	err = service.branchItemRepository.UpdateBranchItem(newBranchItem)
	if err != nil {
		err = errors.New("data branch item failed to be updated")
	}

	return
}

func (service *branchItemService) DeleteBranchItem(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetBranchItem(ctx)
	if err != nil {
		err = errors.New("data branch item not found")
		return
	}

	err = service.branchItemRepository.DeleteBranchItem(id)
	if err != nil {
		err = errors.New("data branch item failed to be deleted")
	}

	return
}

func validateBranchItemReqAndConvertToBranchItem(ctx *gin.Context) (branchItems models.BranchItem, err error) {
	var branchItemsRequest models.BranchItemRequest

	err = ctx.ShouldBindJSON(&branchItemsRequest)
	if err != nil {
		err = errors.New("parameter is not valid")
		return
	}

	err = branchItemsRequest.Validate()
	if err != nil {
		return
	}
	branchItems = branchItemsRequest.ConvertToBranchItem()

	return
}

func validateBranchItemUpdateAndConvertToBranchItem(ctx *gin.Context) (branchItems models.BranchItem, err error) {
	var branchItemsRequest models.BranchItemUpdateRequest

	err = ctx.ShouldBindJSON(&branchItemsRequest)
	if err != nil {
		err = errors.New("parameter is not valid")
		return
	}

	err = branchItemsRequest.Validate()
	if err != nil {
		return
	}
	branchItems = branchItemsRequest.ConvertToBranchItem()

	return
}
