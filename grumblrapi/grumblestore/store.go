package grumblestore

import (
	"fmt"
	"grumblrapi/grumble"

	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

type GrumbleStorer interface {
	GetAll() ([]grumble.Grumble, error)
	Get(id string) (*grumble.Grumble, error)
	Query(querystr string) (*grumble.Grumble, error)
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
func (store *GrumbleStore) GetAll() ([]grumble.Grumble, error) {
	grumbles := make([]grumble.Grumble, 0)
	queryResult, err := store.Scope.Query("SELECT grumbles.* from grumblr.dev.grumbles", nil)
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
	store.Logger.Info(fmt.Sprintf("Successfully queried and recieved %d results", len(grumbles)))
	return grumbles, nil
}

func (store *GrumbleStore) Get(id string) (*grumble.Grumble, error) {
	return nil, nil
}

func (store *GrumbleStore) Query(querystr string) (*grumble.Grumble, error) {
	return nil, nil
}

func (store *GrumbleStore) Insert(id string, grumble *grumble.Grumble) error {
	return nil
}
