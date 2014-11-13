package models

import (
	"time"
)

type Collection struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Albums   []Album   `json:"albums"`
}
