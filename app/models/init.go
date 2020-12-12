package models

import (
	"github.com/revel/revel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	revel.OnAppStart(InitDB)
}

var DB *gorm.DB

func InitDB() {
	dsn, found := revel.Config.String("db.info")
	if !found {
		revel.AppLog.Fatal("Could not find dbInfo")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		revel.AppLog.Fatal("Could not connect to database with error: " + err.Error())
	}
	db.DB()
	db.AutoMigrate(&Service{}, &Product{})
	DB = db
}
