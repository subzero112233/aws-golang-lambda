package users

import "aws-golang-lambda/entity"

type Repository interface {
	GetUser(key string) (entity.User, error)
	AddUser(item entity.User) (err error)
}
