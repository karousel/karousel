package models

import (
	"time"
)

type Collection struct {
	Id       int64     `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
	Albums   []Album   `json:"albums" db:"albums"`
}
