package models

import (
	"time"
)

type Album struct {
	Id       string    `json:"id"`
	Name     string    `json:"name" gorethink:"name"`
	Created  time.Time `json:"created" gorethink:"created"`
	Modified time.Time `json:"modified" gorethink:"modified"`
	Photos   []Photo   `json:"photos" gorethink:"photos"`
}
