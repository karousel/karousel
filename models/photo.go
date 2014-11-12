package models

import (
	"time"
)

type Photo struct {
	Id       string
	Name     string
	Uploaded time.Time
	EXIF     map[string]string
}
