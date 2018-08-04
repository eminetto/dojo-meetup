package model

import "errors"

type User struct {
	Name string
	Age  int
}

func (u *User) IsValid() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Age < 19 {
		return errors.New("invalid age")
	}

	return nil
}

func NewUser(name string, age int) (*User, error) {
	user := User{
		Name: name,
		Age:  age,
	}

	if err := user.IsValid(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) SaveUser() error {
	return nil
}
