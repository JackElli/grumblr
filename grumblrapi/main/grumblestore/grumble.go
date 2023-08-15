package grumblestore

import (
	"time"

	"github.com/google/uuid"
)

type Type string
type DataType string

const (
	Friends Type = "friends"
	Global  Type = "global"

	Image DataType = "image"
	Text  DataType = "text"
)

type Grumble struct {
	Id                string          `json:"id"`
	CreatedBy         string          `json:"createdBy"`
	CreatedByUsername string          `json:"createdByUsername,omitempty"`
	Message           string          `json:"message"`
	Comments          []Comment       `json:"comments"`
	Agrees            map[string]bool `json:"agrees"`
	Disagrees         map[string]bool `json:"disagrees"`
	Date              time.Time       `json:"dateCreated"`
	Type              Type            `json:"type"`
	DataType          DataType        `json:"dataType"`
	Category          string          `json:"category"`
}

func NewGrumble(createdBy string, dateType DataType, message string, _type Type, category string) *Grumble {
	return &Grumble{
		Id:        uuid.New().String(),
		CreatedBy: createdBy,
		Message:   message,
		Comments:  []Comment{},
		Agrees:    make(map[string]bool),
		Disagrees: make(map[string]bool),
		Date:      time.Now(),
		Type:      _type,
		DataType:  dateType,
		Category:  category,
	}
}
