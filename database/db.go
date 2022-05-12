package db

import (
	"MINI_PROJECT_ALTERRA/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var err error

func Init() {
	conf := config.GetConfig()

	DSN := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("connection error!")
	}

	// err = db.Ping()
	// if err != nil {
	// 	panic("DSN invalid")
	// }

}

func CreateCon() *gorm.db {
	return db
}
