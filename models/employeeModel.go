package models

import (
	"errors"
	"time"
)

type Employee struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BranchID  int       `json:"branch_id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EmployeeRequest struct {
	Name     string `json:"name"`
	BranchID int    `json:"branch_id"`
}

type EmployeeOnBranchResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TopEmployeeResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BranchID    int    `json:"branch_id"`
	BranchName  string `json:"branch_name"`
	TotalSales  int    `json:"total_sales"`
	TotalProfit int    `json:"total_profit"`
}

func (e *EmployeeRequest) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	} else if e.BranchID == 0 {
		return errors.New("branch_id is required")
	}
	return nil
}
