package repository

import "go-echo/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User) error
	CreateUser(user *model.User) error
}
