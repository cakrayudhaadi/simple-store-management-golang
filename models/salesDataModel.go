package models

import (
	"errors"
	"simple-store-management/commons"
	"time"
)

type SalesData struct {
	ID         int       `json:"id"`
	BranchID   int       `json:"branch_id"`
	ItemID     int       `json:"item_id"`
	EmployeeID int       `json:"employee_id"`
	Amount     int       `json:"amount"`
	SoldDate   time.Time `json:"sold_date"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (SalesData) TableName() string {
	return "sales_data"
}

type SalesDataRequest struct {
	ItemID     int    `json:"item_id"`
	EmployeeID int    `json:"employee_id"`
	Amount     int    `json:"amount"`
	SoldDate   string `json:"sold_date"`
}

func (s *SalesDataRequest) Validate() error {
	if commons.IsValueEmpty(s.ItemID) {
		return errors.New("item_id is required")
	} else if commons.IsValueEmpty(s.EmployeeID) {
		return errors.New("employee_id is required")
	} else if s.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	} else if commons.IsValueEmpty(s.SoldDate) {
		return errors.New("sold_date is required")
	}
	return nil
}

func (s *SalesDataRequest) ConvertToSalesData() (SalesData, error) {
	soldDate, err := time.Parse("2006-01-02", s.SoldDate)
	return SalesData{
		ItemID:     s.ItemID,
		EmployeeID: s.EmployeeID,
		Amount:     s.Amount,
		SoldDate:   soldDate,
	}, err
}
