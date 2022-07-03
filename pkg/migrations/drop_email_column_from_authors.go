package migrations

import (
	"bookstore-crud/pkg/models"

	"gorm.io/gorm"
)

func dropEmailColumnFromAuthors(migrator gorm.Migrator) error {
	return migrator.DropColumn(&models.Author{}, "email")
}