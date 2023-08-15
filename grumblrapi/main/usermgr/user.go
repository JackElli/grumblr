package usermgr

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id            string    `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	DateCreated   time.Time `json:"dateCreated"`
	Friends       []User    `json:"friends"`
	WelcomePrompt bool      `json:"welcome"`
	Categories    []string  `json:"categories"`
}

func NewUser(username string, password string) *User {
	return &User{
		Id:            uuid.New().String(),
		Username:      username,
		Password:      password,
		DateCreated:   time.Now(),
		Friends:       []User{},
		WelcomePrompt: true,
		Categories:    []string{},
	}
}
