package models

import (
	"time"
)

type User struct {
	Id       int64     `json:"id"`
	Admin    bool      `json:"admin"`
	Username string    `json:"username"`
	Password string    `json:"password,omitempty"`
	Email    string    `json:"email,omitempty"`
	Gravatar string    `json:"gravatar"`
	Joined   time.Time `json:"joined"`
}
