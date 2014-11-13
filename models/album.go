package models

import (
	"time"
)

type Album struct {
	Id           int64     `json:"id"`
	CollectionId int64     `json:"collection_id"`
	Name         string    `json:"name"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Photos       []Photo   `json:"photos"`
}
