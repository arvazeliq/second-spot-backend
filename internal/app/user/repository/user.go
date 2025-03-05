package repository

import (
	"second-spot-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepositoryItf interface {
	Create(user entity.User) (entity.User, error)
	FindByIdentifier(identifier string) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryItf {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepository) FindByIdentifier(identifier string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	return user, err
}
