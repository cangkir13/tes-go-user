package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		panic("failed load env file")
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	fmt.Println(dns)
	gormConfig := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(dns), gormConfig)

	if err != nil {
		return nil, err
	}

	return db, nil
}
