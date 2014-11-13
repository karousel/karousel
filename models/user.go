package models

import (
	"time"
)

type User struct {
	Id       int64     `json:"id" db:"id"`
	Admin    bool      `json:"admin" db:"admin"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password,omitempty" db:"password"`
	Email    string    `json:"email"  db:"email"`
	Gravatar string    `json:"gravatar" db:"gravatar"`
	Joined   time.Time `json:"joined" db:"joined"`
}
