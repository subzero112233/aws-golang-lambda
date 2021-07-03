package handler

import "aws-golang-lambda/entity"

type ErrorOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type HelloOutput struct {
	Message string `json:"message"`
}

type SigninInput struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SigninOutput struct {
	Token string `json:"token"`
}

type SignupInput entity.User

type SignupOutput struct {
	Message string `json:"message"`
}
