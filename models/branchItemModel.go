package models

import "time"

type BranchItem struct {
	ID        int       `json:"id"`
	BranchID  int       `json:"branch_id"`
	ItemID    int       `json:"item_id"`
	Stock     int       `json:"stock"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}
