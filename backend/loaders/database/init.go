package database

import (
	"backend/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func inti() (err error) {
	dsn := config.C.MySql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	return err
}
