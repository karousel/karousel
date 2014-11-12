package models

import (
	"time"
)

type User struct {
	Id       string
	Admin    bool
	Username string
	Password string
	Email    string
	Gravatar string
	Joined   time.Time
}
