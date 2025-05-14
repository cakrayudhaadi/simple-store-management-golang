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

type ItemTypeService interface {
	CreateItemType(ctx *gin.Context) (err error)
	GetAllItemType(ctx *gin.Context) (itemTypes []models.ItemType, err error)
	GetItemType(ctx *gin.Context) (itemType models.ItemType, err error)
	UpdateItemType(ctx *gin.Context) (err error)
	DeleteItemType(ctx *gin.Context) (err error)
}

type itemTypeService struct {
	itemTypeRepository repositories.ItemTypeRepository
}

func NewItemTypeService(itemTypeRepository repositories.ItemTypeRepository) ItemTypeService {
	return &itemTypeService{
		itemTypeRepository,
	}
}

func (service *itemTypeService) CreateItemType(ctx *gin.Context) (err error) {
	var newItemType models.ItemType

	newItemType, err = validateItemTypeReqAndConvertToItemType(ctx)
	if err != nil {
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newItemType.CreatedBy = loginName
	newItemType.CreatedAt = time.Now()

	err = service.itemTypeRepository.CreateItemType(newItemType)
	if err != nil {
		err = errors.New("data itemType gagal dibuat")
	}

	return
}

func (service *itemTypeService) GetAllItemType(ctx *gin.Context) (itemTypes []models.ItemType, err error) {
	itemTypes, err = service.itemTypeRepository.GetAllItemTypes()
	if err != nil {
		err = errors.New("data itemType gagal diambil")
	} else if len(itemTypes) == 0 {
		err = errors.New("data itemType kosong")
	}

	return
}

func (service *itemTypeService) GetItemType(ctx *gin.Context) (itemType models.ItemType, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	itemType, err = service.itemTypeRepository.GetItemType(id)
	if itemType.ID == 0 {
		err = errors.New("data itemType tidak ada")
	} else if err != nil {
		err = errors.New("data itemType gagal diambil")
	}

	return
}

func (service *itemTypeService) UpdateItemType(ctx *gin.Context) (err error) {
	var newItemType models.ItemType
	id, _ := strconv.Atoi(ctx.Param("id"))

	newItemType, err = validateItemTypeReqAndConvertToItemType(ctx)
	if err != nil {
		return
	}

	oldItemType, err := service.GetItemType(ctx)
	if err != nil {
		err = errors.New("data itemType tidak ditemukan")
		return
	}
	newItemType.ID = id
	newItemType.CreatedBy = oldItemType.CreatedBy
	newItemType.CreatedAt = oldItemType.CreatedAt

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newItemType.UpdatedBy = loginName
	newItemType.UpdatedAt = time.Now()

	err = service.itemTypeRepository.UpdateItemType(newItemType)
	if err != nil {
		err = errors.New("data itemType gagal diubah")
	}

	return
}

func (service *itemTypeService) DeleteItemType(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetItemType(ctx)
	if err != nil {
		err = errors.New("data itemType tidak ditemukan")
		return
	}

	err = service.itemTypeRepository.DeleteItemType(id)
	if err != nil {
		err = errors.New("data itemType gagal dihapus")
	}

	return
}

func validateItemTypeReqAndConvertToItemType(ctx *gin.Context) (itemTypes models.ItemType, err error) {
	var itemTypesRequest models.ItemTypeRequest

	err = ctx.ShouldBindJSON(&itemTypesRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = itemTypesRequest.Validate()
	if err != nil {
		return
	}
	itemTypes = itemTypesRequest.ConvertToItemType()

	return
}
