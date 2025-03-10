package bootstrap

import (
	"fmt"
	"second-spot-backend/internal/app/user/interface/rest"
	"second-spot-backend/internal/app/user/repository"
	"second-spot-backend/internal/app/user/usecase"
	"second-spot-backend/internal/infra/bcrypt"
	"second-spot-backend/internal/infra/env"
	"second-spot-backend/internal/infra/fiber"
	"second-spot-backend/internal/infra/jwt"
	"second-spot-backend/internal/infra/mysql"
	"second-spot-backend/internal/infra/validate"
)

func Start() error {
	_env, err := env.New()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_env.DBUsername,
		_env.DBPassword,
		_env.DBHost,
		_env.DBPort,
		_env.DBName,
	)
	db, err := mysql.New(dsn)
	if err != nil {
		return err
	}

	if err := mysql.Migrate(db); err != nil {
		return err
	}

	app := fiber.New()
	validator := validate.New()
	jwtService := jwt.New(_env.JWTSecret)
	bcryptService := bcrypt.New()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, bcryptService, jwtService)
	userHandler := rest.NewUserHandler(userUsecase, validator)

	api := app.Group("/api")
	userHandler.SetupRoutes(api)

	return app.Listen(fmt.Sprintf(":%d", _env.AppPort))
}
