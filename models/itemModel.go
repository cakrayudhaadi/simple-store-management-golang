package models

import (
	"errors"
	"time"
)

type Item struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	ItemTypeID int       `json:"item_type_id"`
	Price      int       `json:"price"`
	Profit     int       `json:"profit"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ItemRequest struct {
	Name       string `json:"name"`
	ItemTypeID int    `json:"item_type_id"`
	Price      int    `json:"price"`
	Profit     int    `json:"profit"`
}

type ItemsOnBranchResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (i *ItemRequest) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	} else if i.ItemTypeID == 0 {
		return errors.New("item_type_id is required")
	} else if i.Price <= 0 {
		return errors.New("price must be greater than 0")
	} else if i.Profit < 0 {
		return errors.New("profit must be greater than or equal to 0")
	} else if i.Profit > i.Price {
		return errors.New("profit must be less than or equal to price")
	}
	return nil
}
