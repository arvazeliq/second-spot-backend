package repository

import (
	"second-spot-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepositoryItf interface {
<<<<<<< HEAD
	Register(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
=======
	Create(user entity.User) (entity.User, error)
	FindByIdentifier(identifier string) (entity.User, error)
>>>>>>> bb86e19 (commit add generate token login)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryItf {
	return &UserRepository{db: db}
}

<<<<<<< HEAD
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
=======
func (r *UserRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepository) FindByIdentifier(identifier string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	return user, err
>>>>>>> bb86e19 (commit add generate token login)
}
