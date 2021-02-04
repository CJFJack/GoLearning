package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	fmt.Println(string(hash), err)

	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("123123")))
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("123456")))
}
