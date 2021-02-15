package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

	var db *sql.DB
	var err error

	stringConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true", user, password, host, port, dbname)
	for {
		db, err = sql.Open("mysql", stringConn)
		if err != nil {
			log.Println(err)
		}

		err = db.Ping()

		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			db.Close()
		} else {
			break
		}
	}

	return db
}
