package models

import (
	"time"

	"gorm.io/gorm"
)

type Jadwal struct {
	Id           int        `gorm:"column:id" json:"id"`
	Day          string     `gorm:"column:hari" json:"hari"`
	Time         string     `gorm:"column:waktu" json:"time"`
	IdMatakuliah int        `json:"-"`
	NidDosen     int        `json:"-"`
	IdRuangan    int        `json:"-"`
	Matakuliah   Matakuliah `json:"matakuliah" gorm:"foreignKey:IdMatakuliah;references:ID"`
	Dosen        Dosen      `json:"dosen" gorm:"foreignKey:NidDosen;references:NID"`
	Ruangan      Ruangan    `json:"ruangan" gorm:"foreignKey:IdRuangan;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Jadwal) TableName() string {
	return "jadwal"
}

func (Matakuliah) TableName() string {
	return "matakuliah"
}

func (Dosen) TableName() string {
	return "dosen"
}

func (Ruangan) TableName() string {
	return "ruangan"
}
