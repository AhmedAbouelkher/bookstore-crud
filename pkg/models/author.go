package models

import "time"

type Author struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
	Books []Book `json:"books" gorm:"constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at,omitempty"`
  	UpdatedAt time.Time `json:"updated_at,omitempty"`
}