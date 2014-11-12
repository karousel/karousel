package models

import (
	"time"
)

type Collection struct {
	Id       string    `json:"id"`
	Name     string    `json:"name" gorethink:"name"`
	Created  time.Time `json:"created" gorethink:"created"`
	Modified time.Time `json:"modified" gorethink:"modified"`
	Albums   []Album   `json:"albums" gorethink:"albums"`
}
