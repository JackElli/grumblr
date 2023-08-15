package userstore

type UserStoreMock struct{}

func NewUserStoreMock() *UserStoreMock {
	return &UserStoreMock{}
}

func (store *UserStoreMock) Get(id string) (*User, error) {
	switch id {
	case "test1":
		return &User{
			Id:       "test1",
			Username: "test1",
			Password: "test",
		}, nil
	case "test2":
		return &User{
			Id:       "test2",
			Username: "test2",
			Password: "test3",
		}, nil
	}
	return nil, nil
}

func (store *UserStoreMock) GetByUsername(username string) (*User, error) {
	return nil, nil
}

func (store *UserStoreMock) Update(id string, user *User) error {
	return nil
}

func (store *UserStoreMock) Insert(id string, user *User) error {
	return nil
}
