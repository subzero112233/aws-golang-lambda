package entity

import "errors"

var ErrUserAlreadyExists = errors.New("Username already exists")

var ErrUserDoesNotExist = errors.New("Username does not exist")

var ErrInvalidPassword = errors.New("Invalid password")

var ErrInvalidInput = errors.New("Invalid input")
