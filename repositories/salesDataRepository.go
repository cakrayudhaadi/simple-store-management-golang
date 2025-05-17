package repositories

import (
	"errors"
	"fmt"
	"simple-store-management/models"
	"time"

	"gorm.io/gorm"
)

type SalesDataRepository interface {
	GetAllSalesDatas() (salesDatas []models.SalesData, err error)
	GetSalesData(id int) (salesData models.SalesData, err error)
	CreateSalesData(salesData models.SalesData) error
	UpdateSalesData(salesData models.SalesData) error
	DeleteSalesData(id int) error
	GetSalesDataBranch(month, year, branchId int) (salesData []models.SalesDataResponse, err error)
	GetSalesDataEmployee(month, year, employeeId int) (salesData []models.SalesDataResponse, err error)
}

type salesDataRepository struct {
	db *gorm.DB
}

func NewSalesDataRepository(database *gorm.DB) SalesDataRepository {
	return &salesDataRepository{
		db: database,
	}
}

func (repo *salesDataRepository) GetAllSalesDatas() (salesDatas []models.SalesData, err error) {
	err = repo.db.Find(&salesDatas).Error
	return
}

func (repo *salesDataRepository) GetSalesData(id int) (salesData models.SalesData, err error) {
	err = repo.db.Where("id = ?", id).Find(&salesData).Error
	if salesData.ID == 0 {
		err = errors.New("sales data not found")
	}
	return
}

func (repo *salesDataRepository) CreateSalesData(salesData models.SalesData) (err error) {
	err = repo.db.Create(&salesData).Error
	return
}

func (repo *salesDataRepository) UpdateSalesData(salesData models.SalesData) (err error) {
	err = repo.db.Model(&models.SalesData{}).Where("id = ?", salesData.ID).Updates(salesData).Error
	return
}

func (repo *salesDataRepository) DeleteSalesData(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&models.SalesData{}).Error
	return
}

func (repo *salesDataRepository) GetSalesDataBranch(month, year, branchId int) (salesData []models.SalesDataResponse, err error) {
	var timestampStart string
	var timestampEnd string
	if year == 0 && month == 0 {
		err = repo.db.Table("sales_data").
			Select("branch.name AS branch_name, employee.name AS employee_name, item.name AS item_name, amount, sold_date").
			Joins("JOIN branch ON sales_data.branch_id = branch.id").
			Joins("JOIN employee ON sales_data.employee_id = employee.id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.branch_id = ?", branchId).
			Scan(&salesData).Error
		if err != nil {
			return
		}
	} else if month == 0 {
		timestampStart = fmt.Sprintf("%d-01-01 00:00:00", year)
		timestampEnd = fmt.Sprintf("%d-12-31 23:59:59", year)

		err = repo.db.Table("sales_data").
			Select("branch.name AS branch_name, employee.name AS employee_name, item.name AS item_name, amount, sold_date").
			Joins("JOIN branch ON sales_data.branch_id = branch.id").
			Joins("JOIN employee ON sales_data.employee_id = employee.id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.sold_date >= ? AND sales_data.sold_date <= ? AND sales_data.branch_id = ?", timestampStart, timestampEnd, branchId).
			Scan(&salesData).Error
		if err != nil {
			return
		}
	} else {
		firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		timestampStart = fmt.Sprintf("%d-%d-01 00:00:00", year, month)
		timestampEnd = fmt.Sprintf("%d-%d-%d 23:59:59", year, month, lastOfMonth.Day())

		err = repo.db.Table("sales_data").
			Select("branch.name AS branch_name, employee.name AS employee_name, item.name AS item_name, amount, sold_date").
			Joins("JOIN branch ON sales_data.branch_id = branch.id").
			Joins("JOIN employee ON sales_data.employee_id = employee.id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.sold_date >= ? AND sales_data.sold_date <= ? AND sales_data.branch_id = ?", timestampStart, timestampEnd, branchId).
			Scan(&salesData).Error
		if err != nil {
			return
		}
	}
	return
}

func (repo *salesDataRepository) GetSalesDataEmployee(month, year, employeeId int) (salesData []models.SalesDataResponse, err error) {
	var timestampStart string
	var timestampEnd string
	if year == 0 && month == 0 {
		err = repo.db.Table("sales_data").
			Select("branch.name AS branch_name, employee.name AS employee_name, item.name AS item_name, amount, sold_date").
			Joins("JOIN branch ON sales_data.branch_id = branch.id").
			Joins("JOIN employee ON sales_data.employee_id = employee.id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.employee_id = ?", employeeId).
			Scan(&salesData).Error
		if err != nil {
			return
		}
	} else if month == 0 {
		timestampStart = fmt.Sprintf("%d-01-01 00:00:00", year)
		timestampEnd = fmt.Sprintf("%d-12-31 23:59:59", year)

		err = repo.db.Table("sales_data").
			Select("branch.name AS branch_name, employee.name AS employee_name, item.name AS item_name, amount, sold_date").
			Joins("JOIN branch ON sales_data.branch_id = branch.id").
			Joins("JOIN employee ON sales_data.employee_id = employee.id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.sold_date >= ? AND sales_data.sold_date <= ? AND sales_data.employee_id = ?", timestampStart, timestampEnd, employeeId).
			Scan(&salesData).Error
		if err != nil {
			return
		}
	} else {
		firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		timestampStart = fmt.Sprintf("%d-%d-01 00:00:00", year, month)
		timestampEnd = fmt.Sprintf("%d-%d-%d 23:59:59", year, month, lastOfMonth.Day())

		err = repo.db.Table("sales_data").
			Select("branch.name AS branch_name, employee.name AS employee_name, item.name AS item_name, amount, sold_date").
			Joins("JOIN branch ON sales_data.branch_id = branch.id").
			Joins("JOIN employee ON sales_data.employee_id = employee.id").
			Joins("JOIN item ON sales_data.item_id = item.id").
			Where("sales_data.sold_date >= ? AND sales_data.sold_date <= ? AND sales_data.employee_id = ?", timestampStart, timestampEnd, employeeId).
			Scan(&salesData).Error
		if err != nil {
			return
		}
	}
	return
}
