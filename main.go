package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	password2 := "secret"
	hash, _ := HashPassword(password)

	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)

	match := CheckPasswordHash(password2, hash)
	fmt.Println("Match: ", match)
}
