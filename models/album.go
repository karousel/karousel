package models

import (
	"time"
)

type Album struct {
	Id       string
	Name     string
	Created  time.Time
	Modified time.Time
	Photos   []Photo
}
