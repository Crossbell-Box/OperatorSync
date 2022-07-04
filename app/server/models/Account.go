package models

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	// Database related fields
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	types.Account
}
