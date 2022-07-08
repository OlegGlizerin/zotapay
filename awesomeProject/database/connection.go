package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@/public"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
}
