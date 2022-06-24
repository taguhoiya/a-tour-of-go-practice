package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"qiita/model"
)

var (
	DbEngine *gorm.DB
)

func init() {
	conninfo := "user=postgres password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}

	dbName := "qiita_dev"
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		//handle the error
		dbname := "dbname=qiita_dev"
		dsn := fmt.Sprintf("%v %v", conninfo, dbname)
		DbEngine, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil && err.Error() != "" {
			log.Fatal(err.Error())
		}
		// // Migrate the schema
		DbEngine.AutoMigrate(&model.Book{})
	}
}
