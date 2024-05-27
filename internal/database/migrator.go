package database

import (
	"go-privfile/internal/domain/shortner"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&shortner.Shortner{})
}
