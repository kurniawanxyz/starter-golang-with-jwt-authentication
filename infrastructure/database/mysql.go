package database

import (
	"fmt"
	"log"

	"github.com/kurniawanxzy/backend-olshop/config"
	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Load() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ENV.DBUser,
		config.ENV.DBPass,
		config.ENV.DBHost,
		config.ENV.DBPort,
		config.ENV.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.TokenVerification{})

	DB = db
	fmt.Println("Database connection established")
}