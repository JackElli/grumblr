package grumble

import "time"

type Grumble struct {
	CreatedBy string    `json:"createdBy"`
	Message   string    `json:"message"`
	Date      time.Time `json:"dateCreated"`
}

func NewGrumble(createdBy string, message string, date time.Time) *Grumble {
	return &Grumble{
		CreatedBy: createdBy,
		Message:   message,
		Date:      date,
	}
}
