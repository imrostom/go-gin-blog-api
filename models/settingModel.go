package models

import (
	"time"
)

type Setting struct {
	ID        uint   `gorm:"primarykey"`
	Key       string `gorm:"type:varchar(255);unique;not null" json:"key"`
	Value     string `gorm:"type:text;not null" json:"value"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
