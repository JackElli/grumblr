package grumblestore

import (
	"grumblrapi/main/grumble"
	"strings"
	"time"
)

type GrumbleStoreMock struct{}

func NewGrumbleStoreMock() *GrumbleStoreMock {
	return &GrumbleStoreMock{}
}

func (store *GrumbleStoreMock) GetAll() ([]grumble.Grumble, error) {
	grumbles := []grumble.Grumble{
		{

			CreatedBy: "user:1",
			Message:   "This is a grumble",
			Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
			Type:      grumble.Global,
		},
		{

			CreatedBy: "user:2",
			Message:   "This is another grumble",
			Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
			Type:      grumble.Friends,
		},
	}
	return grumbles, nil
}

func (store *GrumbleStoreMock) Get(id string) (*grumble.Grumble, error) {
	return nil, nil
}

func (store *GrumbleStoreMock) Query(querystr string) ([]grumble.Grumble, error) {
	if strings.Contains(querystr, "type='friends'") {
		return []grumble.Grumble{
			{

				CreatedBy: "user:1",
				Message:   "This is a friends grumble",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      grumble.Friends,
			},
			{

				CreatedBy: "user:2",
				Message:   "This is another grumble",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      grumble.Friends,
			},
		}, nil
	}

	if strings.Contains(querystr, "type='global'") {
		return []grumble.Grumble{
			{

				CreatedBy: "user:1",
				Message:   "This is a global grumble",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      grumble.Global,
			},
			{

				CreatedBy: "user:2",
				Message:   "This is another global grumble, very public",
				Date:      time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
				Type:      grumble.Global,
			},
		}, nil
	}

	return nil, nil
}

func (store *GrumbleStoreMock) Insert(id string, grumble *grumble.Grumble) error {
	return nil
}
