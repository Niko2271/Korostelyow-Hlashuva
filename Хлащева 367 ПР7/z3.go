package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type User struct {
	Username string
	Email    string
	Password string 
}


func (us *User) SetPassword(password string) {
	hash := sha256.Sum256([]byte(password))
	us.Password = hex.EncodeToString(hash[:])
}


func (us *User) VerifyPassword(password string) bool {
	hash := sha256.Sum256([]byte(password))
	return us.Password == hex.EncodeToString(hash[:])
}

func main() {

user := User{
Username: "alice",
Email:    "alice@example.com",
}

user.SetPassword("mypassword123")
fmt.Printf("Хэш пароля: %s\n", user.Password)

fmt.Printf("Проверка правильного пароля: %v\n", user.VerifyPassword("mypassword123"))
fmt.Printf("Проверка неправильного пароля: %v\n", user.VerifyPassword("wrongpassword"))

user2 := User{
Username: "bob", 
Email:    "bob@example.com",
}
user2.SetPassword("mypassword123")
fmt.Printf("\nХэш того же пароля у другого пользователя: %s\n", user2.Password)
fmt.Printf("Хэши совпадают: %v\n", user.Password == user2.Password)
}
