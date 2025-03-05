package env

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Env struct {
	AppPort int `env:"APP_PORT"`

	DBUsername string `env:"DB_USER"`
	DBPassword string `env:"DB_PASS"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`
<<<<<<< HEAD
=======

	JWTSecret string `env:"JWT_SECRET"`
>>>>>>> bb86e19 (commit add generate token login)
}

func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := new(Env)
	err = env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
