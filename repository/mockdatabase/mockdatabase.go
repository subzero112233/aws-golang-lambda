package mockdatabase

import (
	"aws-golang-lambda/entity"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type MockDatabase struct{}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{}
}

func (r *MockDatabase) AddUser(details entity.User) error {
	users := []string{"existing-dummy-user"}
	for _, a := range users {
		if a == details.Username {
			return fmt.Errorf("user already exists")
		}
	}
	return nil
}

func (r *MockDatabase) GetUser(username string) (details entity.User, err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("encrypted-dummy-pass"), bcrypt.MinCost)
	existingUser := entity.User{
		Address:   "fake-address",
		FirstName: "foo",
		LastName:  "batz",
		Password:  string(hash),
		Username:  "existing-dummy-user",
	}

	if username == existingUser.Username {
		return existingUser, nil
	}

	return details, nil
}
