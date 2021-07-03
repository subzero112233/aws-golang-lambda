package handler

import (
	"aws-golang-lambda/api/middleware"
	"aws-golang-lambda/entity"
	"aws-golang-lambda/usecase/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	Usecase users.UseCase
}

func NewGinHandler(usecase users.UseCase, jwtSecret string) *gin.Engine {
	h := &GinHandler{
		Usecase: usecase,
	}

	r := gin.Default()
	r.POST("/users/signin", h.signIn)
	r.POST("/users/signup", h.signUp)
	r.GET("/hello", middleware.TokenAuthMiddleware(jwtSecret), h.sayHello)
	return r
}

func (h *GinHandler) signIn(c *gin.Context) {
	var l SigninInput

	err := Unmarshal(c, &l)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	err = l.Validate()
	if err != nil {
		ErrHandler(entity.ErrInvalidInput, c)
		return
	}

	var user entity.User
	user.Username, user.Password = l.Username, l.Password

	token, err := h.Usecase.SignIn(user)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	res := SigninOutput{Token: token}

	c.JSON(200, res)
}

func (h *GinHandler) signUp(c *gin.Context) {
	var l entity.User

	err := Unmarshal(c, &l)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	err = l.Validate()
	if err != nil {
		ErrHandler(entity.ErrInvalidInput, c)
		return
	}

	err = h.Usecase.SignUp(l)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	res := SignupOutput{Message: fmt.Sprintf("user %s created successfully", l.Username)}
	c.JSON(201, res)
}

func (h *GinHandler) sayHello(c *gin.Context) {
	var l entity.User
	l.Username = c.GetString("username")

	message, err := h.Usecase.SayHello(l)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	res := HelloOutput{Message: message}

	c.JSON(200, res)
}
