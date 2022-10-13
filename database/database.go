package database

import (
	"fmt"
	"log"
	"notcoder_exe__blog/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Connect with database
var Db *gorm.DB
var err error

func Connect() {
	dsn := "sqlserver://sa:root@localhost:1433?database=notcoderexedb"
	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	Db.AutoMigrate(&models.UserType{}, &models.User{}, &models.Field{}, &models.Tags{}, &models.Post{})
}
