package models

import (
	"time"
)

type Photo struct {
	Id       string            `json:"id"`
	Name     string            `json:"name" gorethink:"name"`
	Uploaded time.Time         `json:"uploaded" gorethink:"uploaded"`
	EXIF     map[string]string `json:"exif" gorethink:"exif"`
}
