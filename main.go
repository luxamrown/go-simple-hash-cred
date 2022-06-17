package main

import (
	"fmt"

	"mohamadelabror.com/gohashcred/repository"
)

func main() {
	passRepo := repository.NewPassword()

	password := "@Bulungan2018"
	password2 := "123"
	err := passRepo.SavePassword(password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Password %s saved \n", password)

	fmt.Printf("Comparing with %s \n", password2)

	isMatch := passRepo.CheckPassword(password2)

	if isMatch == false {
		fmt.Printf("%s and %s is not same", password, password2)
		return
	}
	fmt.Printf("%s and %s is ", password, password2)

}
