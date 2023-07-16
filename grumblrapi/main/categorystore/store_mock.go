package categorystore

import (
	"strings"
)

type CategoryStoreMock struct{}

func NewCategoryStoreMock() *CategoryStoreMock {
	return &CategoryStoreMock{}
}

func (store *CategoryStoreMock) Get(id string) (*Category, error) {
	return nil, nil
}

func (store *CategoryStoreMock) Query(querystr string) ([]Category, error) {
	if strings.Contains(querystr, "type='friends'") {
		return []Category{
			{
				Id:   "testcat1",
				Type: "friends",
				People: []string{
					"jack",
				},
				Name: "Weather",
			},
		}, nil
	}

	return nil, nil
}

func (store *CategoryStoreMock) Insert(id string, grumble *Category) error {
	return nil
}
