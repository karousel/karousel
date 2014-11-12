package models

import (
	"time"
)

type Album struct {
	Name     string
	Created  time.Time
	Modified time.Time
	Photos   []Photo
}
