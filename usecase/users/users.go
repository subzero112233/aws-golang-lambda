package users

import (
	"aws-golang-lambda/entity"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (ser *Service) SignUp(input entity.User) error {
	// need to hash the password
	hashed, err := HashAndSalt(input.Password)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to encrypt password with error: %s", err)
		return err
	}

	details, err := ser.Repository.GetUser(input.Username)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to get user with error: %s", err)
		return err
	}

	if details.Username == input.Username {
		return entity.ErrUserAlreadyExists
	}

	// sending the input but only changing the password to the hashed
	input.Password = hashed

	err = ser.Repository.AddUser(input)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to add user with error: %s", err)
		return err
	}
	return nil
}

func (ser *Service) SignIn(input entity.User) (token string, err error) {
	details, err := ser.Repository.GetUser(input.Username)
	if err != nil {
		return token, err
	}

	if details.Username == "" {
		return token, entity.ErrUserDoesNotExist
	}

	if !(ComparePasswords(details.Password, input.Password)) {
		return token, entity.ErrInvalidPassword
	}

	token, err = CreateToken(details.Username)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to create token with error: %s", err)
		return token, err
	}
	return token, nil
}

func (ser *Service) SayHello(input entity.User) (message string, err error) {
	details, err := ser.Repository.GetUser(input.Username)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to retrieve user data: %s", err)
		return message, err
	}

	message = fmt.Sprintf("Hello %s", details.FirstName)

	return message, nil
}

func ComparePasswords(hashed, plain string) bool {
	// compare the hashed and plain passwords
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func CreateToken(userName string) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	now := time.Now().Local()
	token.Claims = jwt.MapClaims{
		"username": userName,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Hour * time.Duration(1)).Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("t0k3n"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashAndSalt(password string) (string, error) {
	// hash and salt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
