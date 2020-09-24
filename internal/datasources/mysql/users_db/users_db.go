package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mySQLUsername = "MYSQL_USER"
	mySQLPassword = "MYSQL_PASSWORD"
	mySQLHost     = "MYSQL_HOST"
	mySQLSchema   = "MYSQL_DATABASE"
)

var (
	Client *sql.DB

	username = os.Getenv(mySQLUsername)
	password = os.Getenv(mySQLPassword)
	host     = os.Getenv(mySQLHost)
	schema   = os.Getenv(mySQLSchema)
)

// init establish the connection with the MySQL database.
func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	Client = db
	log.Println("database successfully configured")
}
