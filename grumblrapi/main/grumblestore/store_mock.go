package grumblestore

import (
	"strings"
	"time"
)

type GrumbleStoreMock struct{}

func NewGrumbleStoreMock() *GrumbleStoreMock {
	return &GrumbleStoreMock{}
}

func (store *GrumbleStoreMock) Get(id string) (*Grumble, error) {
	return nil, nil
}

func (store *GrumbleStoreMock) Query(querystr string) ([]Grumble, error) {
	if strings.Contains(querystr, "type='friends'") {
		return []Grumble{
			{

				CreatedBy: "user:1",
				Message:   "This is a friends grumble",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      Friends,
			},
			{

				CreatedBy: "user:2",
				Message:   "This is another grumble",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      Friends,
			},
		}, nil
	}

	if strings.Contains(querystr, "type='global'") {
		return []Grumble{
			{

				CreatedBy: "user:1",
				Message:   "This is a global grumble",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      Global,
			},
			{

				CreatedBy: "user:2",
				Message:   "This is another global grumble, very public",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      Global,
			},
		}, nil
	}

	return nil, nil
}

func (store *GrumbleStoreMock) Insert(id string, grumble *Grumble) error {
	return nil
}

func (store *GrumbleStoreMock) Update(id string, grumble *Grumble) error {
	return nil
}
