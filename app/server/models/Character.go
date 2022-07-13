package models

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"gorm.io/gorm"
	"time"
)

type Character struct {
	// Database related fields
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	types.Character
}
