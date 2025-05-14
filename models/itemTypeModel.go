package models

import (
	"errors"
	"time"
)

type ItemType struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemTypeRequest struct {
	Type string `json:"type"`
}

func (i *ItemTypeRequest) Validate() error {
	if i.Type == "" {
		return errors.New("type is required")
	}
	return nil
}
