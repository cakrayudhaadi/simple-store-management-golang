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

type SalesDataService interface {
	CreateSalesData(ctx *gin.Context) (err error)
	GetAllSalesData(ctx *gin.Context) (salesDatas []models.SalesData, err error)
	GetSalesData(ctx *gin.Context) (salesData models.SalesData, err error)
	UpdateSalesData(ctx *gin.Context) (err error)
	DeleteSalesData(ctx *gin.Context) (err error)
}

type salesDataService struct {
	salesDataRepository  repositories.SalesDataRepository
	branchItemRepository repositories.BranchItemRepository
}

func NewSalesDataService(salesDataRepository repositories.SalesDataRepository,
	branchItemRepository repositories.BranchItemRepository) SalesDataService {
	return &salesDataService{
		salesDataRepository,
		branchItemRepository,
	}
}

func (service *salesDataService) CreateSalesData(ctx *gin.Context) (err error) {
	var newSalesData models.SalesData

	newSalesData, err = validateSalesDataReqAndConvertToSalesData(ctx)
	if err != nil {
		return
	}

	branchItem, err := service.branchItemRepository.GetBranchItemByBranchIdAndItemId(newSalesData.BranchID, newSalesData.ItemID)
	if err != nil {
		err = errors.New("branch item tidak ada")
		return
	}

	if branchItem.Stock < newSalesData.Amount {
		err = errors.New("stock tidak cukup")
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newSalesData.CreatedBy = loginName
	newSalesData.CreatedAt = time.Now()
	branchItem.Stock -= newSalesData.Amount

	err = service.salesDataRepository.CreateSalesData(newSalesData)
	if err != nil {
		err = errors.New("data salesData gagal dibuat")
		return
	}

	err = service.branchItemRepository.UpdateBranchItem(branchItem)
	if err != nil {
		err = errors.New("data salesData gagal dibuat")
	}

	return
}

func (service *salesDataService) GetAllSalesData(ctx *gin.Context) (salesDatas []models.SalesData, err error) {
	salesDatas, err = service.salesDataRepository.GetAllSalesDatas()
	if err != nil {
		err = errors.New("data salesData gagal diambil")
	} else if len(salesDatas) == 0 {
		err = errors.New("data salesData kosong")
	}

	return
}

func (service *salesDataService) GetSalesData(ctx *gin.Context) (salesData models.SalesData, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	salesData, err = service.salesDataRepository.GetSalesData(id)
	if salesData.ID == 0 {
		err = errors.New("data salesData tidak ada")
	} else if err != nil {
		err = errors.New("data salesData gagal diambil")
	}

	return
}

func (service *salesDataService) UpdateSalesData(ctx *gin.Context) (err error) {
	var newSalesData models.SalesData
	id, _ := strconv.Atoi(ctx.Param("id"))

	newSalesData, err = validateSalesDataReqAndConvertToSalesData(ctx)
	if err != nil {
		return
	}

	oldSalesData, err := service.GetSalesData(ctx)
	if err != nil {
		err = errors.New("data salesData tidak ditemukan")
		return
	}
	newSalesData.ID = id
	newSalesData.CreatedBy = oldSalesData.CreatedBy
	newSalesData.CreatedAt = oldSalesData.CreatedAt

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newSalesData.UpdatedBy = loginName
	newSalesData.UpdatedAt = time.Now()

	err = service.salesDataRepository.UpdateSalesData(newSalesData)
	if err != nil {
		err = errors.New("data salesData gagal diubah")
	}

	return
}

func (service *salesDataService) DeleteSalesData(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetSalesData(ctx)
	if err != nil {
		err = errors.New("data salesData tidak ditemukan")
		return
	}

	err = service.salesDataRepository.DeleteSalesData(id)
	if err != nil {
		err = errors.New("data salesData gagal dihapus")
	}

	return
}

func validateSalesDataReqAndConvertToSalesData(ctx *gin.Context) (salesDatas models.SalesData, err error) {
	var salesDatasRequest models.SalesDataRequest

	err = ctx.ShouldBindJSON(&salesDatasRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = salesDatasRequest.Validate()
	if err != nil {
		return
	}
	salesDatas = salesDatasRequest.ConvertToSalesData()

	return
}
