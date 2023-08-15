package userstore

import (
	"fmt"

	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

var col = "users"

type UserStorer interface {
	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(id string, user *User) error
	Insert(id string, user *User) error
}

type UserStore struct {
	Logger     *zap.Logger
	Collection *gocb.Collection
	Scope      *gocb.Scope
}

func NewUserStore(logger *zap.Logger, scope *gocb.Scope) *UserStore {
	return &UserStore{
		Logger:     logger,
		Collection: scope.Collection(col),
		Scope:      scope,
	}
}

// Get returns a user based on an id
func (store *UserStore) Get(id string) (*User, error) {
	userResult, err := store.Collection.Get(id, nil)
	if err != nil {
		return nil, err
	}

	var user User
	err = userResult.Content(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByUsername returns a user based on a username
func (store *UserStore) GetByUsername(username string) (*User, error) {
	userResult, err := store.Scope.Query(
		fmt.Sprintf("SELECT users.* FROM grumblr.dev.users WHERE username='%s'", username),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var user User
	err = userResult.One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert inserts a user into the db
func (store *UserStore) Insert(id string, user *User) error {
	_, err := store.Collection.Insert(id, *user, nil)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a users information
func (store *UserStore) Update(id string, user *User) error {
	_, err := store.Collection.Upsert(id, *user, nil)
	if err != nil {
		return err
	}
	return nil
}
