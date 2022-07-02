package models

import (
	"time"
)

type Book struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Price float32  `json:"price"`
	CreatedAt time.Time `json:"created_at"`
  	UpdatedAt time.Time `json:"updated_at"`
}