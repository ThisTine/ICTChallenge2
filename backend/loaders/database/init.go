package database

import (
	"backend/types/database"
	"backend/utils/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (err error) {
	dsn := config.C.MySql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db

	// Initialize model migrations
	// should have mySqlMigration (bool) in config.yaml?
	if true {
		if err := migrate(); err != nil {
			logrus.Fatal("UNABLE TO MIGRATE GORM MODEL")
		}
	}
	assignModel()
	logrus.Debugln("INITIALIZE MYSQL CONNECTION")

	return err
}

func migrate() error {
	if err := DB.AutoMigrate(
		new(database.Card),
		new(database.Raw),
		new(database.Team),
		new(database.Score),
		new(database.Topic),
	); err != nil {
		return err
	}
	return nil
}
