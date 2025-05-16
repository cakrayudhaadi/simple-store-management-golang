package models

import (
	"errors"
	"simple-store-management/commons"
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

type TopEmployeeRequest struct {
	BranchID int `json:"branch_id"`
	Month    int `json:"month"`
	Year     int `json:"year"`
}

type TopEmployeeResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BranchID    int    `json:"branch_id"`
	BranchName  string `json:"branch_name"`
	TotalSales  int    `json:"total_sales"`
	TotalProfit int    `json:"total_profit"`
}

func (Employee) TableName() string {
	return "employee"
}

func (e *EmployeeRequest) Validate() error {
	if commons.IsValueEmpty(e.Name) {
		return errors.New("name is required")
	} else if commons.IsValueEmpty(e.BranchID) {
		return errors.New("branch_id is required")
	}
	return nil
}

func (e *EmployeeRequest) ConvertToEmployee() Employee {
	return Employee{
		Name:     e.Name,
		BranchID: e.BranchID,
	}
}
