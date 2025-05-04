package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func MigrateUpDB(db *sql.DB) error {
	// Create the storage_providers table if it doesn't exist
	_, err := db.Exec(`
	CREATE TYPE storage_provider_file_system AS ENUM (
	'UNIX',
	'WINDOWS'
	);
	
	CREATE TABLE IF NOT EXISTS storage_providers (
			id uuid PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			file_system storage_provider_file_system NOT NULL,
			protocol_connection JSONB NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func MigrateDropDB(db *sql.DB) error {
	// Drop the storage_providers table if it exists
	_, err := db.Exec(`
	DROP TABLE IF EXISTS storage_providers;
	DROP TYPE IF EXISTS storage_provider_file_system;
	`)
	if err != nil {
		return err
	}

	return nil
}

func OpenDB() (*sql.DB, error) {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/ftransfer?sslmode=disable")

	if err != nil {
		return nil, err
	}

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// TODO move this to a config
	// if err := MigrateDropDB(db); err != nil {
	// 	return nil, err
	// }
	// Migrate the database
	// if err := MigrateUpDB(db); err != nil {
	// 	return nil, err
	// }

	// Return the database connection
	return db, nil
}
