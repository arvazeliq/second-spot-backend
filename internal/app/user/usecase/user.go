package usecase

import (
	"errors"
	"second-spot-backend/internal/app/user/repository"
	"second-spot-backend/internal/domain/dto"
	"second-spot-backend/internal/domain/entity"
	"second-spot-backend/internal/infra/bcrypt"
	"second-spot-backend/internal/infra/validate"

	"github.com/google/uuid"
)

type UserUsecaseItf interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type UserUsecase struct {
	userRepo repository.UserRepositoryItf
	validate *validate.Validator
}

func NewUserUsecase(userRepo repository.UserRepositoryItf, validate *validate.Validator) UserUsecaseItf {
	return &UserUsecase{
		userRepo: userRepo,
		validate: validate,
	}
}

func (u *UserUsecase) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if err := u.validate.StructValidate(req); err != nil {
		return nil, errors.New("invalid input data: " + err.Error())
	}

	existingUser, _ := u.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &entity.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := u.userRepo.Register(user); err != nil {
		return nil, errors.New("failed to register user: " + err.Error())
	}

	// Return response DTO
	return &dto.RegisterResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
