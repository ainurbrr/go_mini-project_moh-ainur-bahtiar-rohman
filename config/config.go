package config

import (
	"fmt"

	models "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	InitDB()
	InitialMigration()
	return DB
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "penggalangan-dana",
	}

	// config := Config{
	// 	DB_Username: "u1606266_database",
	// 	DB_Password: "bMZRKgzTX5EwUcm",
	// 	DB_Port:     "3306",
	// 	DB_Host:     "45.130.230.51",
	// 	DB_Name:     "u1606266_crowdfunding",
	// }

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
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
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Campaign_image{})
	DB.AutoMigrate(&models.Campaign{})
	DB.AutoMigrate(&models.Transaction{})
}
