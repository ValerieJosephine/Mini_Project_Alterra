package repository

import (
	"MINI_PROJECT_ALTERRA/models"
	// "gorm.io/gorm"
)

// type UserRepository interface {
// 	GetUsers() ([]models.UserCostum, error)
// }

// type URepository struct {
// 	db *gorm.DB
// }

// func NewUserRepository(db *gorm.DB) *URepository {
// 	return &URepository{db: db}

// }

func (r *Repository) GetUsers() ([]models.User, error) {
	users := []models.User{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil

}
