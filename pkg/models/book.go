package models

import (
	"time"
)

type Book struct {
	ID uint `gorm:"primaryKey" json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Price float32  `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
  	UpdatedAt time.Time `json:"updated_at,omitempty"`
}