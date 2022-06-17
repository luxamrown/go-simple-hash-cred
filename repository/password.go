package repository

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordRepo interface {
	SavePassword(password string) error
	CheckPassword(password string) bool
}

type passwordRepo struct {
	pass string
}

func (p *passwordRepo) SavePassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	p.pass = string(bytes)
	return nil
}

func (p *passwordRepo) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.pass), []byte(password))
	return err == nil
}

func NewPassword() PasswordRepo {
	return &passwordRepo{}
}
