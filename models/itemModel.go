package models

import "time"

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
