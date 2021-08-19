package migrations

import (
	"github.com/ygorrodriguesmeli/gorestapi/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Product{})
}
