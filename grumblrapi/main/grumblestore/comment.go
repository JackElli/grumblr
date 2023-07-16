package grumblestore

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        string    `json:"id"`
	CreatedBy string    `json:"createdBy"`
	Message   string    `json:"message"`
	Date      time.Time `json:"dateCreated"`
}

func NewComment(createdBy string, message string) *Comment {
	return &Comment{
		Id:        uuid.New().String(),
		CreatedBy: createdBy,
		Message:   message,
		Date:      time.Now(),
	}
}
