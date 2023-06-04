package grumble

import "time"

type Type string

const (
	Friends Type = "friends"
	Global  Type = "global"
)

type Grumble struct {
	Id        string    `json:"id"`
	CreatedBy string    `json:"createdBy"`
	Message   string    `json:"message"`
	Date      time.Time `json:"dateCreated"`
	Type      Type      `json:"type"`
}

func NewGrumble(createdBy string, message string, date time.Time, _type Type) *Grumble {
	return &Grumble{
		CreatedBy: createdBy,
		Message:   message,
		Date:      date,
		Type:      _type,
	}
}
