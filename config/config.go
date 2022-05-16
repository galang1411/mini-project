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
		DB_Host:       "db",
		DB_Name:       "mini_p",
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
	DB.AutoMigrate(&models.Ruangan{})
}

type configTest struct {
	DB_User_Test string
	DB_Pass_Test string
	DB_Host_Test string
	DB_Port_Test string
	DB_Name_Test string
}

func InitDBTest() {
	config := configTest{
		DB_User_Test: "root",
		DB_Pass_Test: "gromizk123",
		DB_Host_Test: "localhost",
		DB_Port_Test: "3306",
		DB_Name_Test: "unit_test",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_User_Test, config.DB_Pass_Test, config.DB_Host_Test, config.DB_Port_Test, config.DB_Name_Test)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.AutoMigrate(&models.Dosen{})
	DB.AutoMigrate(&models.Matakuliah{})
	DB.AutoMigrate(&models.Ruangan{})
	DB.AutoMigrate(&models.Operator{})
	DB.AutoMigrate(&models.InsertJadwal{})
}
