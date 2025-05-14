package repositories

import (
	"simple-store-management/models"

	"gorm.io/gorm"
)

type SalesDataRepository interface {
	GetAllSalesDatas() (salesDatas []models.SalesData, err error)
	GetSalesData(id int) (salesData models.SalesData, err error)
	CreateSalesData(salesData models.SalesData) error
	UpdateSalesData(salesData models.SalesData) error
	DeleteSalesData(id int) error
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
