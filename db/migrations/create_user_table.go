package migrations

import "gorm.io/gorm"

func CreateUsersTable(db *gorm.DB) error {
	type User struct {
		ID    uint   `gorm:"primaryKey"`
		Name  string `gorm:"size:100"`
		Score int    `gorm:"default:0"`
	}

	return db.AutoMigrate(&User{})
}
