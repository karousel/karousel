package models

import (
	"time"
)

type Collection struct {
	Id       string
	Name     string
	Created  time.Time
	Modified time.Time
	Albums   []Album
}
