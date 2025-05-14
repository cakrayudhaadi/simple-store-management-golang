package models

import (
	"errors"
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

type SalesDataRequest struct {
	BranchID   int       `json:"branch_id"`
	ItemID     int       `json:"item_id"`
	EmployeeID int       `json:"employee_id"`
	Amount     int       `json:"amount"`
	SoldDate   time.Time `json:"sold_date"`
}

func (s *SalesDataRequest) Validate() error {
	if s.BranchID == 0 {
		return errors.New("branch_id is required")
	} else if s.ItemID == 0 {
		return errors.New("item_id is required")
	} else if s.EmployeeID == 0 {
		return errors.New("employee_id is required")
	} else if s.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	} else if s.SoldDate.IsZero() {
		return errors.New("sold_date is required")
	}
	return nil
}

func (s *SalesDataRequest) ConvertToSalesData() SalesData {
	return SalesData{
		BranchID:   s.BranchID,
		ItemID:     s.ItemID,
		EmployeeID: s.EmployeeID,
		Amount:     s.Amount,
		SoldDate:   s.SoldDate,
	}
}
