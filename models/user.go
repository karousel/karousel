package models

import (
	"time"
)

type User struct {
	Admin    bool
	Username string
	Password string
	Email    string
	Gravatar string
	Joined   time.Time
}
