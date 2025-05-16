package services

import (
	"errors"
	"simple-store-management/models"
	"simple-store-management/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	CreateItem(ctx *gin.Context) (err error)
	GetAllItem(ctx *gin.Context) (items []models.Item, err error)
	GetItem(ctx *gin.Context) (item models.Item, err error)
	UpdateItem(ctx *gin.Context) (err error)
	DeleteItem(ctx *gin.Context) (err error)
}

type itemService struct {
	itemRepository     repositories.ItemRepository
	itemTypeRepository repositories.ItemTypeRepository
}

func NewItemService(itemRepository repositories.ItemRepository,
	itemTypeRepository repositories.ItemTypeRepository,
) ItemService {
	return &itemService{
		itemRepository,
		itemTypeRepository,
	}
}

func (service *itemService) CreateItem(ctx *gin.Context) (err error) {
	var newItem models.Item

	newItem, err = validateItemReqAndConvertToItem(ctx)
	if err != nil {
		return
	}

	_, err = service.itemTypeRepository.GetItemType(newItem.ItemTypeID)
	if err != nil {
		err = errors.New("data item type not found")
		return
	}

	// loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	// newItem.CreatedBy = loginName
	newItem.CreatedAt = time.Now()

	err = service.itemRepository.CreateItem(newItem)
	if err != nil {
		err = errors.New("data item failed to be created")
	}

	return
}

func (service *itemService) GetAllItem(ctx *gin.Context) (items []models.Item, err error) {
	items, err = service.itemRepository.GetAllItems()
	if err != nil {
		err = errors.New("data item failed to be loaded")
	} else if len(items) == 0 {
		err = errors.New("data item kosong")
	}

	return
}

func (service *itemService) GetItem(ctx *gin.Context) (item models.Item, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	item, err = service.itemRepository.GetItem(id)

	return
}

func (service *itemService) UpdateItem(ctx *gin.Context) (err error) {
	var newItem models.Item
	id, _ := strconv.Atoi(ctx.Param("id"))

	newItem, err = validateItemReqAndConvertToItem(ctx)
	if err != nil {
		return
	}

	oldItem, err := service.GetItem(ctx)
	if err != nil {
		err = errors.New("data item not found")
		return
	}

	_, err = service.itemTypeRepository.GetItemType(newItem.ItemTypeID)
	if err != nil {
		err = errors.New("data item type not found")
		return
	}

	newItem.ID = id
	newItem.CreatedBy = oldItem.CreatedBy
	newItem.CreatedAt = oldItem.CreatedAt

	// loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	// newItem.UpdatedBy = loginName
	newItem.UpdatedAt = time.Now()

	err = service.itemRepository.UpdateItem(newItem)
	if err != nil {
		err = errors.New("data item failed to be updated")
	}

	return
}

func (service *itemService) DeleteItem(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetItem(ctx)
	if err != nil {
		err = errors.New("data item not found")
		return
	}

	err = service.itemRepository.DeleteItem(id)
	if err != nil {
		err = errors.New("data item failed to be deleted")
	}

	return
}

func validateItemReqAndConvertToItem(ctx *gin.Context) (items models.Item, err error) {
	var itemsRequest models.ItemRequest

	err = ctx.ShouldBindJSON(&itemsRequest)
	if err != nil {
		err = errors.New("parameter is not valid")
		return
	}

	err = itemsRequest.Validate()
	if err != nil {
		return
	}
	items = itemsRequest.ConvertToItem()

	return
}
