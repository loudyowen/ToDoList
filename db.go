package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		panic(err)
	}
	database := os.Getenv("MYSQL_DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS activities (
			id INT(11) NOT NULL AUTO_INCREMENT,
			title VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todo_items (
		id INT(11) NOT NULL AUTO_INCREMENT,
		activity_group_id INT(11) NOT NULL,
		title VARCHAR(255) NOT NULL,
		is_active TINYINT(1) NOT NULL,
		priority VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
