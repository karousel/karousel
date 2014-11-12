package models

import (
	"time"
)

type User struct {
	Id       string    `json:"id"`
	Admin    bool      `json:"admin" gorethink:"admin"`
	Username string    `json:"username" gorethink:"username"`
	Password string    `json:"-" gorethink:"password"`
	Email    string    `json:"email" gorethink:"email"`
	Gravatar string    `json:"gravatar" gorethink:"gravatar"`
	Joined   time.Time `json:"joined" gorethink:"joined"`
}
