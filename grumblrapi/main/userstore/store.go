package userstore

import (
	"grumblrapi/main/user"

	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

var col = "users"

type UserStorer interface {
	Get(id string) (*user.User, error)
	Insert(id string, user *user.User) error
}

type UserStore struct {
	Logger     *zap.Logger
	Collection *gocb.Collection
}

func NewUserStore(logger *zap.Logger, scope *gocb.Scope) *UserStore {
	return &UserStore{
		Logger:     logger,
		Collection: scope.Collection(col),
	}
}

// Get returns a user based on an id
func (store *UserStore) Get(id string) (*user.User, error) {
	userResult, err := store.Collection.Get(id, nil)
	if err != nil {
		return nil, err
	}

	var user user.User
	err = userResult.Content(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert inserts a user into the db
func (store *UserStore) Insert(id string, user *user.User) error {
	_, err := store.Collection.Insert(id, *user, nil)
	if err != nil {
		return err
	}
	return nil
}
