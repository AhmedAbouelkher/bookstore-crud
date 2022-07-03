package models

import "time"

type Author struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at,omitempty"`
  	UpdatedAt time.Time `json:"updated_at,omitempty"`
}