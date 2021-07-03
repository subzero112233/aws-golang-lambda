package jwtparser

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestValidateToken(t *testing.T) {
	type args struct {
		secret     string
		username   string
		expiration int64
	}
	test := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "successful token validation",
			args:    args{secret: "s3cr3t", username: "y0zer", expiration: time.Now().Add(time.Hour * time.Duration(1)).Unix()},
			wantErr: false,
		},
		{
			name:    "unsuccessful token validation due to a bad secret",
			args:    args{secret: "bad-secret", username: "y0zer", expiration: time.Now().Add(time.Hour * time.Duration(1)).Unix()},
			wantErr: true,
		},
		{
			name:    "unsuccessful token validation due expired token",
			args:    args{secret: "s3cr3t", username: "y0zer", expiration: time.Now().Add(time.Hour * time.Duration(-1)).Unix()},
			wantErr: true,
		},
	}

	for _, tt := range test {
		t.Run("test ValidateToken with "+tt.name, func(t *testing.T) {

			// create token
			token := jwt.New(jwt.GetSigningMethod("HS256"))

			now := time.Now().Local()
			token.Claims = jwt.MapClaims{
				"username": tt.args.username,
				"iat":      now.Unix(),
				"exp":      tt.args.expiration,
			}

			tokenString, _ := token.SignedString([]byte("s3cr3t"))

			_, err := ValidateToken(tokenString, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing ValidateToken. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
