package usecase

import (
	"errors"
	"second-spot-backend/internal/app/user/repository"
	"second-spot-backend/internal/domain/dto"
	"second-spot-backend/internal/domain/entity"
	"second-spot-backend/internal/infra/bcrypt"
	"second-spot-backend/internal/infra/jwt"

	"github.com/google/uuid"
)

type UserUsecaseItf interface {
	Register(req dto.RegisterRequest) (dto.RegisterResponse, error)
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
}

type UserUsecase struct {
	repo repository.UserRepositoryItf
	hash bcrypt.BcryptItf
	jwt  jwt.JWTItf
}

func NewUserUsecase(repo repository.UserRepositoryItf, hash bcrypt.BcryptItf, jwt jwt.JWTItf) UserUsecaseItf {
	return &UserUsecase{repo: repo, hash: hash, jwt: jwt}
}

func (u *UserUsecase) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	hashedPassword, err := u.hash.HashPassword(req.Password)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	user := entity.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}

	createdUser, err := u.repo.Create(user)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		ID:       createdUser.ID.String(),
		Email:    createdUser.Email,
		Username: createdUser.Username,
	}, nil
}

func (u *UserUsecase) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := u.repo.FindByIdentifier(req.Identifier)
	if err != nil {
		return dto.LoginResponse{}, errors.New("user not found")
	}

	if !u.hash.ComparePassword(user.Password, req.Password) {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := u.jwt.GenerateToken(user.ID.String())
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		ID:    user.ID.String(),
		Token: token,
	}, nil
}
