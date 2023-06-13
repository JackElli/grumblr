package user

import "time"

type User struct {
	Id            string    `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	DateCreated   time.Time `json:"dateCreated"`
	Friends       []string  `json:"friends"`
	WelcomePrompt bool      `json:"welcome"`
}

func NewUser(username string, password string) *User {
	return &User{
		Username:      username,
		Password:      password,
		DateCreated:   time.Now(),
		Friends:       []string{},
		WelcomePrompt: true,
	}
}
