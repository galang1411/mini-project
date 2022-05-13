package models

import (
	"time"

	"gorm.io/gorm"
)

type Matakuliah struct {
	ID        uint           `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:nama" json:"nama"`
	SKS       int            `gorm:"column:sks" json:"sks"`
	Semester  int            `gorm:"column: semester" json:"semester"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
