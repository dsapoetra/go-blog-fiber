package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Encrypt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePass(hashed, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err
}
