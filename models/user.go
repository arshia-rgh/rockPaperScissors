package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:100"`
	Score     int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"default:autoCreateTime"`
}

func NewUser(id uint, name string, score int) *User {
	return &User{
		ID:        id,
		Name:      name,
		Score:     score,
		CreatedAt: time.Now(),
	}

}
