package models

import (
	"time"

	"gorm.io/gorm"
)

type Jadwal struct {
	Id         int    `gorm:"column:id" json:"id"`
	Day        string `gorm:"column:hari" json:"hari"`
	Time       string `gorm:"column:waktu" json:"time"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Matakuliah Matakuliah     `json:"matakuliah" gorm:"foreignKey:IDMatakuliah;references:ID"`
	Dosen      Dosen          `json:"dosen" gorm:"foreignKey:NIDDosen;references:NID"`
	Ruangan    Ruangan        `json:"ruangan" gorm:"foreignKey:IDruangan;references:ID"`
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
