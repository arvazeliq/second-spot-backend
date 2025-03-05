package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

<<<<<<< HEAD
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) bool {
=======
type BcryptItf interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
}

type Bcrypt struct{}

func New() BcryptItf {
	return &Bcrypt{}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (b *Bcrypt) ComparePassword(hashedPassword, password string) bool {
>>>>>>> bb86e19 (commit add generate token login)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
