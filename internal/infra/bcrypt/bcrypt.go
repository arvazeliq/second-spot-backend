package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

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
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
