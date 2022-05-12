package models

import (
	"time"
	
	"gorm.io/gorm"
)

type Dosen struct {
	NID       int   `gorm:"column:nid" json:"nid"`
	Name      string `gorm:"column:nama" json:"nama"`
	Gender    string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Major     string `gorm:"column:jurusan" json:"jurusan"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
