package models

import (
	"time"
)

type Photo struct {
	Name     string
	Uploaded time.Time
	EXIF     map[string]string
}
