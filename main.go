package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/george124816/job-dev-backend-interview/internal/pkg/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	db := db.ConnectDB()
	err := db.Ping()
	if err != nil {
		fmt.Println("Não conectado")
	} else {
		fmt.Println("Conectado")
	}

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
