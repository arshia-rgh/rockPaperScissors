package repositories

import (
	"gorm.io/gorm"
	"rockPaperScissors/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	err := r.DB.First(&user, id).Error

	return &user, err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Updates(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.DB.Delete(id).Error
}
