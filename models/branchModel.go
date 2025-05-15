package models

import (
	"errors"
	"time"
)

type Branch struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BranchRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BranchResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type EmployeesOfBranchResponse struct {
	ID        int                        `json:"id"`
	Name      string                     `json:"name"`
	Address   string                     `json:"address"`
	Employees []EmployeeOnBranchResponse `json:"employees"`
}

type ItemsOfBranchResponse struct {
	ID      int                     `json:"id"`
	Name    string                  `json:"name"`
	Address string                  `json:"address"`
	Items   []ItemsOnBranchResponse `json:"items"`
}

type TopBranchRequest struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

type TopBranchResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	TotalSales  int    `json:"total_sales"`
	TotalProfit int    `json:"total_profit"`
}

func (b *BranchRequest) Validate() error {
	if b.Name == "" {
		return errors.New("name is required")
	} else if b.Address == "" {
		return errors.New("address is required")
	}
	return nil
}

func (b *BranchRequest) ConvertToBranch() Branch {
	return Branch{
		Name:    b.Name,
		Address: b.Address,
	}
}
