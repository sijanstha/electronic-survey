package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMysqlStorage() (*sql.DB, error) {
	dbHost := os.Getenv("DATA_SOURCE_HOST")
	dbPort := os.Getenv("DATA_SOURCE_PORT")
	dbPassword := os.Getenv("DATA_SOURCE_PASSWORD")
	dbUser := os.Getenv("DATA_SOURCE_USER")
	dbName := os.Getenv("DB_NAME")
	migrationURL := os.Getenv("MIGRATION_URL")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	migrationDbSource := "mysql://%s:%s@tcp(%s:%s)/%s"
	runDBMigration(migrationURL, fmt.Sprintf(migrationDbSource, dbUser, dbPassword, dbHost, dbPort, dbName))

	return db, nil
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
