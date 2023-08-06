package userstore

import (
	"grumblrapi/main/user"
)

type UserStoreMock struct{}

func NewUserStoreMock() *UserStoreMock {
	return &UserStoreMock{}
}

func (store *UserStoreMock) Get(id string) (*user.User, error) {
	switch id {
	case "test1":
		return user.NewUser(
			"test",
			"test",
		), nil
	case "test2":
		return user.NewUser(
			"test2",
			"test",
		), nil
	}
	return nil, nil
}

func (store *UserStoreMock) Update(id string, user *user.User) error {
	return nil
}

func (store *UserStoreMock) Insert(id string, user *user.User) error {
	return nil
}
