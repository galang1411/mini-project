package config

import (
	"fmt"
	"mini-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_matkulname string
	DB_Password   string
	DB_Port       string
	DB_Host       string
	DB_Name       string
}

func InitDB() {

	config := Config{
		DB_matkulname: "root",
		DB_Password:   "gromizk123",
		DB_Port:       "3306",
		DB_Host:       "localhost",
		DB_Name:       "mini_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_matkulname,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.Dosen{})
	DB.AutoMigrate(&models.Matakuliah{})
}
