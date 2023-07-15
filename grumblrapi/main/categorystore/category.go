package categorystore

import (
	"github.com/google/uuid"
)

type Category struct {
	Id     string   `json:"id"`
	Type   string   `json:"type"`
	People []string `json:"people"`
	Name   string   `json:"name"`
}

func NewCategory(_type string, people []string, name string) *Category {
	return &Category{
		Id:     uuid.New().String(),
		Type:   _type,
		People: people,
		Name:   name,
	}
}
