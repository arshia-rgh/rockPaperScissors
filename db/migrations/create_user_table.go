package migrations

import (
	"gorm.io/gorm"
	"rockPaperScissors/models"
)

func CreateUsersTable(db *gorm.DB) error {

	return db.AutoMigrate(&models.User{})
}
