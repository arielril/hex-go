package model

import (
	"fmt"
	"hexgo/internal/types"
)

type userModel struct {
	db types.Persistence
}

func NewUser(db types.Persistence) types.UserModel {
	return userModel{db}
}

func (u userModel) CreateUser(user *types.UserModel) {
	fmt.Println("Creating the user", user)
}

func (u userModel) FindUser(params map[string]interface{}) {
	fmt.Println("Trying to find a user with, params")
}
