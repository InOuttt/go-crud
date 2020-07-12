package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GetHashedPassword(password string) string {
	pass := []byte(password)
	hashedPass, errCrypt := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if errCrypt != nil {
		panic(errCrypt)
	}
	return string(hashedPass)
}

func ComparePassword(hashedPass []byte, password string) error {
	pass := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedPass, pass)
	return err
}