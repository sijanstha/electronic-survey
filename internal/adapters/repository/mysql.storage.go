package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlStorage() (*sql.DB, error) {
	dbHost := os.Getenv("DATA_SOURCE_HOST")
	dbPort := os.Getenv("DATA_SOURCE_PORT")
	dbPassword := os.Getenv("DATA_SOURCE_PASSWORD")
	dbUser := os.Getenv("DATA_SOURCE_USER")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
