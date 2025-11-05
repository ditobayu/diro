// Package main provides database migration utilities
// This command-line tool manages database schema migrations using golang-migrate
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// main is the entry point for the migration command
// It provides a command-line interface for managing database migrations
func main() {
	// Define command-line flags for migration operations
	var action string
	flag.StringVar(&action, "action", "", "Migration action: up, down, version, force, create")
	var steps int
	flag.IntVar(&steps, "steps", 0, "Number of steps for up/down migration")
	var version int
	flag.IntVar(&version, "version", 0, "Version for force migration")
	var name string
	flag.StringVar(&name, "name", "", "Migration name for create action")
	flag.Parse()

	// Display usage information if no action is specified
	if action == "" {
		fmt.Println("Usage:")
		fmt.Println("  go run cmd/migrate/migrate.go -action=up                    # Run all up migrations")
		fmt.Println("  go run cmd/migrate/migrate.go -action=up -steps=1           # Run 1 up migration")
		fmt.Println("  go run cmd/migrate/migrate.go -action=down -steps=1         # Run 1 down migration")
		fmt.Println("  go run cmd/migrate/migrate.go -action=version               # Show current version")
		fmt.Println("  go run cmd/migrate/migrate.go -action=force -version=1      # Force version")
		fmt.Println("  go run cmd/migrate/migrate.go -action=create -name=add_table # Create new migration")
		os.Exit(1)
	}

	// Load database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "diro_db")

	// Construct MySQL Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Handle migration file creation separately (doesn't need database connection)
	if action == "create" {
		if name == "" {
			log.Fatal("Migration name is required for create action")
		}
		createMigration(name)
		return
	}

	// Open database connection for migration operations
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Create migrate driver instance for MySQL
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Failed to create driver instance:", err)
	}

	// Create migrate instance with file source and database driver
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Migration files directory
		"mysql",             // Database type
		driver,
	)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}

	// Execute the requested migration action
	switch action {
	case "up":
		// Apply migrations to bring database schema up to latest version
		if steps > 0 {
			err = m.Steps(steps) // Apply specific number of migrations
		} else {
			err = m.Up() // Apply all pending migrations
		}
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Failed to run up migration:", err)
		}
		fmt.Println("Up migration completed successfully")

	case "down":
		// Rollback migrations to previous schema versions
		if steps > 0 {
			err = m.Steps(-steps) // Rollback specific number of migrations
		} else {
			err = m.Down() // Rollback all migrations
		}
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Failed to run down migration:", err)
		}
		fmt.Println("Down migration completed successfully")

	case "version":
		// Display current migration version and dirty state
		currentVersion, dirty, err := m.Version()
		if err != nil {
			log.Fatal("Failed to get version:", err)
		}
		fmt.Printf("Current version: %d, Dirty: %t\n", currentVersion, dirty)

	case "force":
		// Force migration to a specific version (use with caution)
		if version == 0 {
			log.Fatal("Version is required for force action")
		}
		err = m.Force(version)
		if err != nil {
			log.Fatal("Failed to force version:", err)
		}
		fmt.Printf("Forced version to: %d\n", version)

	default:
		log.Fatal("Invalid action. Use: up, down, version, force, or create")
	}
}

// createMigration generates new migration files with proper naming and structure
func createMigration(name string) {
	// Get the next migration number by scanning existing files
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal("Failed to read migrations directory:", err)
	}

	maxVersion := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var version int
		n, err := fmt.Sscanf(file.Name(), "%06d_", &version)
		if err == nil && n == 1 && version > maxVersion {
			maxVersion = version
		}
	}

	nextNumber := maxVersion + 1

	// Create migration file names with sequential numbering
	upFile := fmt.Sprintf("migrations/%06d_%s.up.sql", nextNumber, name)
	downFile := fmt.Sprintf("migrations/%06d_%s.down.sql", nextNumber, name)

	// Create up migration file with template content
	upContent := fmt.Sprintf("-- Migration: %s\n-- Created at: %s\n\n-- Write your up migration here\n", name, "")
	err = os.WriteFile(upFile, []byte(upContent), 0644)
	if err != nil {
		log.Fatal("Failed to create up migration file:", err)
	}

	// Create down migration file with template content
	downContent := fmt.Sprintf("-- Migration: %s\n-- Created at: %s\n\n-- Write your down migration here\n", name, "")
	err = os.WriteFile(downFile, []byte(downContent), 0644)
	if err != nil {
		log.Fatal("Failed to create down migration file:", err)
	}

	fmt.Printf("Created migration files:\n")
	fmt.Printf("  %s\n", upFile)
	fmt.Printf("  %s\n", downFile)
}

// getEnv retrieves an environment variable value or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
