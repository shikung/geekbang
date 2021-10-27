package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	MYSQL_HOST     string = "139.159.190.137"
	MYSQL_PORT     string = "3306"
	MYSQL_USERNAME string = "haiji"
	MYSQL_PASS     string = "haiJi@2020"
	MYSQL_CHARSET  string = "utf8"
	MYSQL_DB       string = "test"
)

var DB *sql.DB

type Test struct {
	id    int
	name  string
	title string
}

func ConncetDB() {
	mysql_connect_uri := MYSQL_USERNAME + ":" + MYSQL_PASS + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DB + "?charset=" + MYSQL_CHARSET
	var err error
	DB, err = sql.Open("mysql", mysql_connect_uri)
	if err != nil {
		panic("unknown driver or db uri format error")
	}
	err = DB.Ping()
	if err != nil {
		panic("db connect error")
	}
}

//
func FetchRowData() (t Test, err error) {
	var test Test
	err1 := DB.QueryRow("SELECT id,name,title FROM test WHERE status=?", 1).Scan(&test.id, &test.name, &test.title)
	switch {
	case err1 == sql.ErrNoRows:
		return test, errors.Wrap(err1, "Fetch Row data is empty")
	case err1 != nil:
		return test, errors.Wrap(err1, "Fetch Row data is error")
	}
	return test, nil
}

func main() {
	ConncetDB()
	test, err := FetchRowData()
	fmt.Print(test)
	fmt.Print(errors.Unwrap(err))

}
