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

type Migrator interface {
	AutoMigrate(dst ...interface{}) error

	CurrentDatabase() string

	CreateTable(dst ...interface{}) error
	DropTable(dst ...interface{}) error
	HasTable(dst interface{}) bool
	RenameTable(oldName,newName interface{}) error

	AddColumn(dst interface{}, field string) error
	DropColumn(dst interface{}, field string) error
	AlterColumn(dst interface{}, field string) error
	HasColumn(dst interface{}, field string) bool
	RenameColumn(dst interface{}, oldName, field string) error
	MigrateColumn(dst interface{}, field *schema.Field, columnType *sql.ColumnType) error
	ColumnTypes(dst interface{}) ([]*sql.ColumnType, error)

	CreateConstraint(dst interface{}, name string) error
	DropConstraint(dst interface{}, name string) error
	HasConstraint(dst interface{}, name string) bool

	CreateIndex(dst interface{}, name string) error
	DropIndex(dst interface{}, name string) error
	HasIndex(dst interface{}, name string) bool
	RenameIndex(dst interface{}, oldName, newName string) error
}


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
