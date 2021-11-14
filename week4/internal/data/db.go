package data

import (
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/pkg/log"
	_ "github.com/go-sql-driver/mysql"
)

// Data .
type Data struct {
	db *sql.DB
}

const (
	MYSQL_HOST     string = "111.111.111.137"
	MYSQL_PORT     string = "111"
	MYSQL_USERNAME string = "1111"
	MYSQL_PASS     string = "111@2020"
	MYSQL_CHARSET  string = "utf8"
	MYSQL_DB       string = "test"
)

func NewDB() (db *sql.DB, err error) {
	mysqlConnectUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s??charset=%s", MYSQL_USERNAME, MYSQL_PASS, MYSQL_HOST, MYSQL_PORT, MYSQL_DB, MYSQL_CHARSET)
	db, err = sql.Open("mysql", mysqlConnectUri)
	if err != nil {
		panic("unknown driver or db uri format error")
	}
	err = db.Ping()
	if err != nil {
		panic("db connect error")
	}
	return db, err
}

func NewData(database *sql.DB) (*Data, func(), error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	d := &Data{
		db:  database
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
