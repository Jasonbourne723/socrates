package models

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	Id int64 `json:"id" gorm:"primaryKey"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
