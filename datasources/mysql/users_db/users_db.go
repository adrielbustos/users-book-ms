package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
// mysql_users_username = "mysql_users_username"
// mysql_users_password = ""
// mysql_users_host     = ""
// mysql_users_schema   = ""
)

var (
	Client *sql.DB

	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// schema   = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"",
		"127.0.0.1:3306",
		"go_course_ms",
	)
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	// mysql.SetLogger("")
	log.Panicln("db connected")
}
