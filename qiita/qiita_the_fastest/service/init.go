package service

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"qiita/model"
	"strings"
)

var (
	DbEngine *gorm.DB
	Dbname   string = "qiita_dev"
	err      error
)

type DbConfig struct {
	user     string
	password string
	host     string
	sslmode  string
}

func config(dbConfig DbConfig) string {
	hoge := map[string]string{"user": dbConfig.user, "password": dbConfig.password, "host": dbConfig.host, "sslmode": dbConfig.sslmode}
	var foo []string
	for k, v := range hoge {
		pair := k + "=" + v
		foo = append(foo, pair)
	}
	str2 := strings.Join(foo, " ")
	return str2
}

func create(dbConfig DbConfig) {
	db, err := sql.Open("postgres", config(dbConfig))
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + Dbname)
	if err != nil {
		return
	}

	_, err = db.Exec("USE " + Dbname)
	if err != nil {
		return
	}
}

func init() {
	dbConfig := DbConfig{"postgres", "postgres", "localhost", "disable"}
	create(dbConfig)
	dsn := config(dbConfig) + " " + "dbname=" + Dbname
	DbEngine, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	// // Migrate the schema
	DbEngine.AutoMigrate(&model.Book{})
}
