package models

import "time"

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
