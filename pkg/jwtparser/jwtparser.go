package jwtparser

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	IssuedAt   int64  `json:"iat"`
	Expiration int64  `json:"exp"`
	Username   string `json:"username"`
}

func ValidateToken(t string, secret string) (claims Claims, err error) {
	// Parse jwt token
	token, err := jwt.ParseWithClaims(
		t,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil || !token.Valid {
		return claims, fmt.Errorf("unable to parse token or it's invalid")
	}

	return claims, nil
}

func (claims *Claims) Valid() error {
	// check whether one of the claims is missing
	if claims.Expiration == 0 || claims.Username == "" {
		return fmt.Errorf("missing jwt fields")
	}

	// check whether the token is expired or not.
	now := time.Now()
	exp := time.Unix(claims.Expiration, 0)
	if now.After(exp) {
		return fmt.Errorf("token is expired")
	}

	return nil
}
