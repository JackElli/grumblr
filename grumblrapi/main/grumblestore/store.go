package grumblestore

import (
	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

var col = "grumbles"

type GrumbleStorer interface {
	Get(id string) (*Grumble, error)
	Query(querystr string) ([]Grumble, error)
	Insert(id string, grumble *Grumble) error
}

type GrumbleStore struct {
	Logger     *zap.Logger
	Scope      *gocb.Scope
	Collection *gocb.Collection
}

func NewGrumbleStore(logger *zap.Logger, scope *gocb.Scope) *GrumbleStore {
	return &GrumbleStore{
		Logger:     logger,
		Scope:      scope,
		Collection: scope.Collection(col),
	}
}

// Get returns a grumble based on an id
func (store *GrumbleStore) Get(id string) (*Grumble, error) {
	return nil, nil
}

// Query allows us to execute a more fine grained query on the scope
func (store *GrumbleStore) Query(querystr string) ([]Grumble, error) {
	grumbles := make([]Grumble, 0)
	queryResult, err := store.Scope.Query(querystr, nil)
	if err != nil {
		return nil, err
	}

	var result Grumble
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
func (store *GrumbleStore) Insert(id string, grumble *Grumble) error {
	_, err := store.Collection.Insert(id, *grumble, nil)
	if err != nil {
		return err
	}
	return nil
}
