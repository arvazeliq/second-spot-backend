package usecase

import (
	"errors"
	"second-spot-backend/internal/app/user/repository"
	"second-spot-backend/internal/domain/dto"
	"second-spot-backend/internal/domain/entity"
	"second-spot-backend/internal/infra/bcrypt"
<<<<<<< HEAD
	"second-spot-backend/internal/infra/validate"
=======
	"second-spot-backend/internal/infra/jwt"
>>>>>>> bb86e19 (commit add generate token login)

	"github.com/google/uuid"
)

type UserUsecaseItf interface {
<<<<<<< HEAD
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
=======
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
>>>>>>> bb86e19 (commit add generate token login)
		ID:       uuid.New(),
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}

<<<<<<< HEAD
	if err := u.userRepo.Register(user); err != nil {
		return nil, errors.New("failed to register user: " + err.Error())
	}

	// Return response DTO
	return &dto.RegisterResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Username: user.Username,
=======
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
>>>>>>> bb86e19 (commit add generate token login)
	}, nil
}
