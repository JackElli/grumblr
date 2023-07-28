package categorystore

import (
	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

var col = "categories"

type CategoryStorer interface {
	Get(id string) (*Category, error)
	Query(querystr string) ([]Category, error)
	Insert(id string, category *Category) error
}

type CategoryStore struct {
	Logger     *zap.Logger
	Scope      *gocb.Scope
	Collection *gocb.Collection
}

func NewCategoryStore(logger *zap.Logger, scope *gocb.Scope) *CategoryStore {
	return &CategoryStore{
		Logger:     logger,
		Scope:      scope,
		Collection: scope.Collection(col),
	}
}

// Get returns a category based on an id
func (store *CategoryStore) Get(id string) (*Category, error) {
	return nil, nil
}

// Query allows us to execute a more fine grained query on the scope
func (store *CategoryStore) Query(querystr string) ([]Category, error) {
	categories := make([]Category, 0)
	queryResult, err := store.Scope.Query(querystr, nil)
	if err != nil {
		return nil, err
	}

	for queryResult.Next() {
		var result Category
		err := queryResult.Row(&result)
		if err != nil {
			store.Logger.Error(err.Error())
		}
		categories = append(categories, result)
	}

	return categories, nil
}

// Insert inserts a category into the db
func (store *CategoryStore) Insert(id string, category *Category) error {
	_, err := store.Collection.Insert(id, *category, nil)
	if err != nil {
		return err
	}
	return nil
}
