package models

import (
	"time"
)

type Collection struct {
	Name     string
	Created  time.Time
	Modified time.Time
	Albums   []Album
}
