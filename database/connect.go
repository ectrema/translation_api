package database

import (
	"fmt"
	"gorm.io/driver/mariadb"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "root:Translation_mysql_root@tcp(localhost:3306)/Translation_api?charset=utf8mb4&parseTime=True&loc=Local"
	//Setup db connection form .env data
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to database")
	MigrateDatabase()
}