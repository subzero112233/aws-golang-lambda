package users

import (
	"aws-golang-lambda/entity"
	"aws-golang-lambda/repository/mockdatabase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func TestSignUp(t *testing.T) {
	type args struct {
		input entity.User
	}

	repository := mockdatabase.NewMockDatabase()
	logger := log.With().Logger()

	ser := LoadService(repository, &logger)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful sign up",
			args: args{input: entity.User{
				Address:   "mapo 4",
				FirstName: "charlie",
				LastName:  "miko",
				Username:  "y0zer",
				Password:  "pass",
			},
			},
			wantErr: false,
		},
		{
			name: "Unsuccessful sign up due to user already exists",
			args: args{input: entity.User{
				Address:   "fake-address",
				FirstName: "foo",
				LastName:  "batz",
				Username:  "existing-dummy-user",
				Password:  "pass",
			},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test SignUp with "+tt.name, func(t *testing.T) {
			err := ser.SignUp(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing SignUp operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	type args struct {
		input entity.User
	}

	repository := mockdatabase.NewMockDatabase()
	logger := log.With().Logger()

	ser := LoadService(repository, &logger)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful sign in",
			args: args{input: entity.User{
				Username: "existing-dummy-user",
				Password: "encrypted-dummy-pass",
			},
			},
			wantErr: false,
		},
		{
			name: "Unsuccessful sign in due to user does not exist",
			args: args{input: entity.User{
				Username: "non-existing-dummy-user",
				Password: "encrypted-dummy-pass",
			},
			},
			wantErr: true,
		},
		{
			name: "Unsuccessful sign in due to incorrect password",
			args: args{input: entity.User{
				Username: "y0zer",
				Password: "incorrect-dummy-pass",
			},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test SignIn with "+tt.name, func(t *testing.T) {
			_, err := ser.SignIn(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing SignIn operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSayHello(t *testing.T) {
	repository := mockdatabase.NewMockDatabase()
	logger := log.With().Logger()

	ser := LoadService(repository, &logger)

	t.Run("test SignIn", func(t *testing.T) {
		c := &gin.Context{}
		c.Set("username", "existing-dummy-user")

		l := entity.User{Username: c.GetString("username")}
		_, err := ser.SayHello(l)
		if err != nil {
			t.Errorf("error while executing SayHello operation. error = %v", err)
		}
	})
}
