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
	branchRepository     repositories.BranchRepository
	itemRepository       repositories.ItemRepository
	employeeRepository   repositories.EmployeeRepository
}

func NewSalesDataService(salesDataRepository repositories.SalesDataRepository,
	branchItemRepository repositories.BranchItemRepository,
	branchRepository repositories.BranchRepository,
	itemRepository repositories.ItemRepository,
	employeeRepository repositories.EmployeeRepository,
) SalesDataService {
	return &salesDataService{
		salesDataRepository,
		branchItemRepository,
		branchRepository,
		itemRepository,
		employeeRepository,
	}
}

func (service *salesDataService) CreateSalesData(ctx *gin.Context) (err error) {
	var newSalesData models.SalesData

	newSalesData, err = validateSalesDataReqAndConvertToSalesData(ctx)
	if err != nil {
		return
	}

	_, err = service.itemRepository.GetItem(newSalesData.ItemID)
	if err != nil {
		err = errors.New("data item not found")
		return
	}

	_, err = service.employeeRepository.GetEmployee(newSalesData.EmployeeID)
	if err != nil {
		err = errors.New("data employee not found")
		return
	}

	branchID, err := service.branchRepository.GetBranchIDByEmployeeID(newSalesData.EmployeeID)
	if err != nil {
		err = errors.New("data employee does not belong to any branch")
		return
	}
	newSalesData.BranchID = branchID

	branchItem, err := service.branchItemRepository.GetBranchItemByBranchIDAndItemID(newSalesData.BranchID, newSalesData.ItemID)
	if err != nil {
		err = errors.New("branch item not found")
		return
	}

	if branchItem.ID == 0 {
		err = errors.New("branch does not have this item")
		return
	}

	if branchItem.Stock < newSalesData.Amount {
		err = errors.New("stock not enough")
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
		err = errors.New("data salesData failed to be created")
		return
	}

	err = service.branchItemRepository.UpdateBranchItem(branchItem)
	if err != nil {
		err = errors.New("data salesData failed to be created")
	}

	return
}

func (service *salesDataService) GetAllSalesData(ctx *gin.Context) (salesDatas []models.SalesData, err error) {
	salesDatas, err = service.salesDataRepository.GetAllSalesDatas()
	if err != nil {
		err = errors.New("data salesData failed to be loaded")
	} else if len(salesDatas) == 0 {
		err = errors.New("data salesData kosong")
	}

	return
}

func (service *salesDataService) GetSalesData(ctx *gin.Context) (salesData models.SalesData, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	salesData, err = service.salesDataRepository.GetSalesData(id)

	return
}

func (service *salesDataService) UpdateSalesData(ctx *gin.Context) (err error) {
	var newSalesData models.SalesData
	id, _ := strconv.Atoi(ctx.Param("id"))

	newSalesData, err = validateSalesDataReqAndConvertToSalesData(ctx)
	if err != nil {
		return
	}

	_, err = service.itemRepository.GetItem(newSalesData.ItemID)
	if err != nil {
		err = errors.New("data item not found")
		return
	}

	_, err = service.employeeRepository.GetEmployee(newSalesData.EmployeeID)
	if err != nil {
		err = errors.New("data employee not found")
		return
	}

	branchID, err := service.branchRepository.GetBranchIDByEmployeeID(newSalesData.EmployeeID)
	if err != nil {
		err = errors.New("data employee does not belong to any branch")
		return
	}
	newSalesData.BranchID = branchID

	branchItem, err := service.branchItemRepository.GetBranchItemByBranchIDAndItemID(newSalesData.BranchID, newSalesData.ItemID)
	if err != nil {
		err = errors.New("branch item not found")
		return
	}

	if branchItem.ID == 0 {
		err = errors.New("branch does not have this item")
		return
	}

	if branchItem.Stock < newSalesData.Amount {
		err = errors.New("stock not enough")
		return
	}

	oldSalesData, err := service.GetSalesData(ctx)
	if err != nil {
		err = errors.New("data salesData not found")
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
		err = errors.New("data salesData failed to be updated")
	}

	return
}

func (service *salesDataService) DeleteSalesData(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetSalesData(ctx)
	if err != nil {
		err = errors.New("data salesData not found")
		return
	}

	err = service.salesDataRepository.DeleteSalesData(id)
	if err != nil {
		err = errors.New("data salesData failed to be deleted")
	}

	return
}

func validateSalesDataReqAndConvertToSalesData(ctx *gin.Context) (salesDatas models.SalesData, err error) {
	var salesDatasRequest models.SalesDataRequest

	err = ctx.ShouldBindJSON(&salesDatasRequest)
	if err != nil {
		err = errors.New("parameter is not valid")
		return
	}

	err = salesDatasRequest.Validate()
	if err != nil {
		return
	}
	salesDatas, err = salesDatasRequest.ConvertToSalesData()
	if err != nil {
		return
	}

	return
}
