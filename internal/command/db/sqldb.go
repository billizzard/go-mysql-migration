package db

import (
	"database/sql"
	"fmt"
	"github.com/billizzard/go-mysql-migration/configs"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDb() {
	log.Println("Connecting to MySQL database...")

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		configs.DB_USER,
		configs.DB_PASS,
		configs.DB_HOST,
		configs.DB_PORT,
		configs.DB_NAME,
	))

	if err != nil {
		log.Fatal("Unable to connect to database. Try run 'init' command. ", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Unable to connect to database. Try run 'init' command. ", err.Error())
	}

	log.Println("Database connected")

	DB = db
}

func ConnectMysql() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/",
		configs.DB_USER,
		configs.DB_PASS,
		configs.DB_HOST,
		configs.DB_PORT,
	))

	if err != nil {
		log.Fatal("Unable to connect to mysql. Error: " + err.Error())
	}

	return db
}
