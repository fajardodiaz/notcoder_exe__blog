package database

import (
	"fmt"
	"log"
	"notcoder_exe__blog/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Connect with database
var db *gorm.DB
var err error

func Connect() {
	dsn := "sqlserver://sa:root@localhost:1433?database=notcoderexedb"
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	db.AutoMigrate(&models.UserType{}, &models.User{}, &models.Field{}, &models.Tags{}, &models.Post{})
}
