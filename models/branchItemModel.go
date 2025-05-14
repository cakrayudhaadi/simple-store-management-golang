package models

import (
	"errors"
	"time"
)

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

type AddBranchItemRequest struct {
	BranchID   int `json:"branch_id"`
	ItemID     int `json:"item_id"`
	AddedStock int `json:"added_stock"`
}

type RemoveBranchItemRequest struct {
	BranchID     int `json:"branch_id"`
	ItemID       int `json:"item_id"`
	RemovedStock int `json:"removed_stock"`
}

func (b *AddBranchItemRequest) Validate() (err error) {
	if b.BranchID == 0 {
		return errors.New("branch_id is required")
	} else if b.ItemID == 0 {
		return errors.New("item_id is required")
	} else if b.AddedStock < 0 {
		return errors.New("added_stock must be greater than or equal to 0")
	}
	return nil
}

func (b *RemoveBranchItemRequest) Validate() (err error) {
	if b.BranchID == 0 {
		return errors.New("branch_id is required")
	} else if b.ItemID == 0 {
		return errors.New("item_id is required")
	} else if b.RemovedStock < 0 {
		return errors.New("removed_stock must be greater than or equal to 0")
	}
	return nil
}
