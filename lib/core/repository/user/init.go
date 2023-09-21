package userrepo

import "gorm.io/gorm"

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}
