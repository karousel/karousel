package models

import (
	"time"
)

type Album struct {
	Id       int64     `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
	Photos   []Photo   `json:"photos" db:"photos"`
}
