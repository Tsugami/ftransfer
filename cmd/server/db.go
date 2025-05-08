package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateUpDB(databaseURL string) error {
	// Obter o caminho absoluto do diretório de migrações
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting working directory: %v", err)
	}
	migrationsPath := filepath.Join(wd, "migrations")
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)

	fmt.Println("Migrations path:", migrationsPath)
	fmt.Println("Migrations URL:", migrationsURL)

	m, err := migrate.New(migrationsURL, databaseURL)

	if err != nil {
		return fmt.Errorf("error creating migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error migrating up: %v", err)
	}

	return nil
}

func OpenDB() (*sql.DB, error) {
	// Open a connection to the database
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	fmt.Println("Try to connect to the database with the following URL:", DATABASE_URL)

	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		return nil, err
	}

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Migrate the database
	if err := MigrateUpDB(DATABASE_URL); err != nil {
		return nil, err
	}

	return db, nil
}
