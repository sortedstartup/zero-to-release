package db

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // Import SQLite driver
)

//go:embed migrations
var migrationFiles embed.FS

func MigrateDB(driver string, dbURL string) error {
	// Create a new migration source instance using the embedded migration files
	_migrationFiles, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		log.Fatalf("Failed to load migration files: %v", err)
	}

	log.Printf("Migrating database: %s", dbURL)

	// Create a new migration instance
	m, err := migrate.NewWithSourceInstance("iofs", _migrationFiles, dbURL)
	if err != nil {
		return fmt.Errorf("failed creating new migration: %w", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed while migrating: %w", err)
	}

	log.Println("Migration completed successfully")
	return nil
}
