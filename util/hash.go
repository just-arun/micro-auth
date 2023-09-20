package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type password struct{}

func Password() password {
	return password{}
}

func (r password) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (r password) Compare(hash, password string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}



