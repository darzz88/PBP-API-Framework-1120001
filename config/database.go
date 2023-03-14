package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:@tcp(localhost:3306)/db_tes"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	DB.AutoMigrate()
}
