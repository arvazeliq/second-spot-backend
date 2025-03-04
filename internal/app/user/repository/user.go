package repository

import (
	"second-spot-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepositoryItf interface {
	Register(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryItf {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
