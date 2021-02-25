package util

import (
	"database/sql"
	"github.com/wzslr321/gorm_postgresql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)
var(
	Db *gorm.DB
)

func InitDB() (){

	var err error

	newLogger := logger.New(
		log.New(os.Stdout,"\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Silent,
			Colorful: true,
		},
	)

	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "user=postgres password=postgres dbname=react_golang port=5432 sslmode=disable TimeZone=Europe/Warsaw",
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic("Failed connect to the database")
	}

	_ = Db.AutoMigrate(
		&models.Product{},
	)
}
