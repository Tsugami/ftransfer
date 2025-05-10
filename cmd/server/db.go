package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Database represents a connection to the database
type Database struct {
	db *sql.DB
}

// NewDatabase creates a new instance of Database
func NewDatabase(url string) *Database {
	return &Database{}
}

// Connect establishes a connection to the database
func (d *Database) Connect(url string) error {
	if d.db != nil {
		d.db.Close()
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection
	if err := db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("failed to verify database connection: %v", err)
	}

	d.db = db
	return nil
}

// GetDB returns the database instance
func (d *Database) GetDB() *sql.DB {
	return d.db
}

// Close closes the database connection
func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

func MigrateUpDB(databaseURL string) error {
	// Get the absolute path of the migrations directory
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

func OpenDB() (*Database, error) {
	// Get the database URL
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	fmt.Println("Trying to connect to the database with the following URL:", DATABASE_URL)

	// Create a new database instance
	db := NewDatabase(DATABASE_URL)

	// Connect to the database
	if err := db.Connect(DATABASE_URL); err != nil {
		return nil, err
	}

	// Run migrations
	if err := MigrateUpDB(DATABASE_URL); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
