package migrations

import "bookstore-crud/pkg/configs"

func RunMigrations() {
	db := configs.GetDB()

	dropEmailColumnFromAuthors(db.Migrator())
}