package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var user string = "root"
var password string = "secret"
var host string = "127.0.0.1"
var port int = 3306
var dbname string = "main"

func main() {

	db := connectDB()
	migration(db)

}

func connectDB() *sql.DB {

	stringConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true", user, password, host, port, dbname)

	db, err := sql.Open("mysql", stringConn)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return db
	}

	return db
}

func migration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"mysql",
		driver,
	)
	if err != nil {
		log.Println(err)
	}

	m.Up()

}
