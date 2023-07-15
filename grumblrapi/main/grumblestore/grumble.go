package grumblestore

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	Friends Type = "friends"
	Global  Type = "global"
)

type Grumble struct {
	Id        string    `json:"id"`
	CreatedBy string    `json:"createdBy"`
	Message   string    `json:"message"`
	Comments  []Comment `json:"comments"`
	Date      time.Time `json:"dateCreated"`
	Type      Type      `json:"type"`
	Category  string    `json:"category"`
}

func NewGrumble(createdBy string, message string, _type Type, category string) *Grumble {
	return &Grumble{
		Id:        uuid.New().String(),
		CreatedBy: createdBy,
		Message:   message,
		Comments:  []Comment{},
		Date:      time.Now(),
		Type:      _type,
		Category:  category,
	}
}
