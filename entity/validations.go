package entity

import "errors"

func (input User) Validate() error {
	if input.Username == "" {
		return errors.New("username must be provided")
	}

	if input.Password == "" {
		return errors.New("password must be provided")
	}

	if input.Address == "" {
		return errors.New("address must be provided")
	}

	if input.FirstName == "" {
		return errors.New("first name must be provided")
	}

	if input.LastName == "" {
		return errors.New("last name must be provided")
	}
	return nil
}
