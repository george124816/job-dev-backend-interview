package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = "secret"
	dbname   = "main"
)

func ConnectDB() *sql.DB {
	stringConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true", user, password, host, port, dbname)

	db, err := sql.Open("mysql", stringConn)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}
