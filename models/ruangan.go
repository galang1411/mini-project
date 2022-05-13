package models

import (
	"time"

	"gorm.io/gorm"
)

type Ruangan struct {
	ID        uint           `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:nama" json:"nama"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
