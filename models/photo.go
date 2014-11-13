package models

import (
	"time"
)

type Photo struct {
	Id       int64             `json:"id" db:"id"`
	Name     string            `json:"name" db:"name"`
	Uploaded time.Time         `json:"uploaded" db:"uploaded"`
	EXIF     map[string]string `json:"exif" db:"exif"`
}
