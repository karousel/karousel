package models

import (
	"time"
)

type Photo struct {
	Id         int64     `json:"id"`
	AlbumId    int64     `json:"album_id"`
	Name       string    `json:"name"`
	Uploaded   time.Time `json:"uploaded"`
	EXIFString string    `json:"-"`
}
