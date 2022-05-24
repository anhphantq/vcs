package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	User_id  uint   `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role_id  uint   `json:"role_id,string" validate:"number,min=1"`
	Password string `json:"password" validate:"min=8"`
}

func main() {

	var user Account
	user.Email = "123@gmail.com"
	user.Role_id = 2
	validate := validator.New()
	fmt.Println(validate.Struct(&user))
}
