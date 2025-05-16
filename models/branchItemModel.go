package models

import (
	"errors"
	"simple-store-management/commons"
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

func (BranchItem) TableName() string {
	return "branch_item"
}

type BranchItemRequest struct {
	BranchID   int `json:"branch_id"`
	ItemID     int `json:"item_id"`
	AddedStock int `json:"added_stock"`
}

type BranchItemUpdateRequest struct {
	AddedStock int `json:"added_stock"`
}

func (b *BranchItemRequest) Validate() (err error) {
	if commons.IsValueEmpty(b.BranchID) {
		return errors.New("branch_id is required")
	} else if commons.IsValueEmpty(b.ItemID == 0) {
		return errors.New("item_id is required")
	} else if b.AddedStock < 0 {
		return errors.New("added_stock must be greater than or equal to 0")
	}
	return nil
}

func (b *BranchItemUpdateRequest) Validate() (err error) {
	if b.AddedStock < 0 {
		return errors.New("added_stock must be greater than or equal to 0")
	}
	return nil
}

func (b *BranchItemRequest) ConvertToBranchItem() BranchItem {
	return BranchItem{
		BranchID: b.BranchID,
		ItemID:   b.ItemID,
		Stock:    b.AddedStock,
	}
}

func (b *BranchItemUpdateRequest) ConvertToBranchItem() BranchItem {
	return BranchItem{
		Stock: b.AddedStock,
	}
}
