package grumblemgr

import (
	"grumblrapi/main/usermgr"
	"log"

	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

var col = "grumbles"

type GrumbleStorer interface {
	Get(id string) (*Grumble, error)
	Query(querystr string, options *gocb.QueryOptions) ([]Grumble, error)
	Insert(id string, grumble *Grumble) error
	Update(id string, grumble *Grumble) error
}

type GrumbleStore struct {
	Logger     *zap.Logger
	Scope      *gocb.Scope
	Collection *gocb.Collection
	UserMgr    usermgr.UserStorer
}

func NewGrumbleStore(logger *zap.Logger, scope *gocb.Scope, userMgr usermgr.UserStorer) *GrumbleStore {
	return &GrumbleStore{
		Logger:     logger,
		Scope:      scope,
		Collection: scope.Collection(col),
		UserMgr:    userMgr, // Added to aid with user queries
	}
}

// Get returns a grumble based on an id
func (store *GrumbleStore) Get(id string) (*Grumble, error) {
	grumbleData, err := store.Collection.Get(id, nil)
	if err != nil {
		return nil, err
	}

	var grumble Grumble
	err = grumbleData.Content(&grumble)
	if err != nil {
		log.Fatal(err)
	}

	// Leave error for now (non blocking)
	// TODO check for 404
	grumbleCreator, err := store.UserMgr.Get(grumble.CreatedBy)
	if err == nil {
		grumble.CreatedByUsername = grumbleCreator.Username
	}

	// Update comments users
	for _, c := range grumble.Comments {
		commentCreator, err := store.UserMgr.Get(c.CreatedBy)
		if err != nil {
			continue
		}
		c.CreatedByUsername = commentCreator.Username
	}

	return &grumble, nil
}

// Query allows us to execute a more fine grained query on the scope
func (store *GrumbleStore) Query(querystr string, options *gocb.QueryOptions) ([]Grumble, error) {
	grumbles := make([]Grumble, 0)
	queryResult, err := store.Scope.Query(querystr, options)
	if err != nil {
		return nil, err
	}

	for queryResult.Next() {
		var result Grumble
		err := queryResult.Row(&result)
		if err != nil {
			store.Logger.Error(err.Error())
		}
		grumbles = append(grumbles, result)
	}

	return grumbles, nil
}

// Insert inserts a grumble into the db
func (store *GrumbleStore) Update(id string, grumble *Grumble) error {
	_, err := store.Collection.Upsert(id, *grumble, nil)
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a grumble into the db
func (store *GrumbleStore) Insert(id string, grumble *Grumble) error {
	_, err := store.Collection.Insert(id, *grumble, nil)
	if err != nil {
		return err
	}
	return nil
}