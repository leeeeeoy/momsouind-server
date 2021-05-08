package db

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DBClient    *sqlx.DB
)

func InitDB() {
	conf := fmt.Sprintf("%s:%s@(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sqlx.Connect("mysql", conf)
	if err != nil {
		log.Fatalln(err)
	}
	db.Ping()
	DBClient = db
	DBClient.Mapper = reflectx.NewMapperTagFunc("db",
		nil,
		func(s string) string {
			return s
		},
	)
	DBClient.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
}
