package grumblestore

import (
	"grumblrapi/main/grumble"

	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

type GrumbleStorer interface {
	GetAll() ([]grumble.Grumble, error)
	Get(id string) (*grumble.Grumble, error)
	Query(querystr string) ([]grumble.Grumble, error)
	Insert(id string, grumble *grumble.Grumble) error
}

type GrumbleStore struct {
	Logger     *zap.Logger
	Scope      *gocb.Scope
	Collection *gocb.Collection
}

func NewGrumbleStore(logger *zap.Logger, scope *gocb.Scope, collection *gocb.Collection) *GrumbleStore {
	return &GrumbleStore{
		Logger:     logger,
		Scope:      scope,
		Collection: collection,
	}
}

// GetAll returns all global grumbles in the database with LIMIT 50
func (store *GrumbleStore) GetAll() ([]grumble.Grumble, error) {
	grumbles := make([]grumble.Grumble, 0)
	queryResult, err := store.Scope.Query("SELECT grumbles.* from grumblr.dev.grumbles WHERE type='global' LIMIT 50", nil)
	if err != nil {
		return nil, err
	}
	var result grumble.Grumble
	for queryResult.Next() {
		err := queryResult.Row(&result)
		if err != nil {
			store.Logger.Error(err.Error())
		}
		grumbles = append(grumbles, result)
	}
	return grumbles, nil
}

// Get returns a grumble based on an id
func (store *GrumbleStore) Get(id string) (*grumble.Grumble, error) {
	return nil, nil
}

// Query allows us to execute a more fine grained query on the scope
func (store *GrumbleStore) Query(querystr string) ([]grumble.Grumble, error) {
	grumbles := make([]grumble.Grumble, 0)
	queryResult, err := store.Scope.Query(querystr, nil)
	if err != nil {
		return nil, err
	}
	var result grumble.Grumble
	for queryResult.Next() {
		err := queryResult.Row(&result)
		if err != nil {
			store.Logger.Error(err.Error())
		}
		grumbles = append(grumbles, result)
	}
	return grumbles, nil
}

// Insert inserts a grumble into the db
func (store *GrumbleStore) Insert(id string, grumble *grumble.Grumble) error {
	_, err := store.Collection.Insert(id, *grumble, nil)
	if err != nil {
		return err
	}
	return nil
}
