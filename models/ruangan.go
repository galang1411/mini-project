package models

import (
	"time"

	"gorm.io/gorm"
)

type Ruangan struct {
	ID        uint           `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:nama" json:"nama"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
