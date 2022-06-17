package repository

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type PasswordRepo interface {
	SavePassword(password string) error
	CheckPassword(password string) bool
}

type passwordRepo struct {
	passDb *sqlx.DB
}

func (p *passwordRepo) SavePassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	_, err = p.passDb.Exec("INSERT INTO password_clone(password) VALUES($1)", string(bytes))
	if err != nil {
		panic(err)
	}
	return nil
}

func (p *passwordRepo) CheckPassword(password string) bool {
	var pass string
	err := p.passDb.Get(&pass, "SELECT * FROM password_clone")
	if err != nil {
		panic(err)
	}
	_, err = p.passDb.Exec("DELETE FROM password_clone WHERE password = $1", pass)
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	return err == nil
}

func NewPassword(db *sqlx.DB) PasswordRepo {
	return &passwordRepo{
		passDb: db,
	}
}
