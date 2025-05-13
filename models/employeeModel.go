package models

import "time"

type Employee struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BranchID  int       `json:"branch_id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}
