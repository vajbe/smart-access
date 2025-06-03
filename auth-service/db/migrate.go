package db

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() {
	createMigrationTable := `
    CREATE TABLE IF NOT EXISTS schema_migrations (
        version TEXT PRIMARY KEY
    );
    `
	if _, err := DB.Exec(createMigrationTable); err != nil {
		log.Fatalf("Failed to create schema_migrations table: %v", err)
	}

	applied := map[string]bool{}
	rows, err := DB.Query("SELECT version FROM schema_migrations")
	if err != nil {
		log.Fatalf("Failed to fetch applied migrations: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v string
		rows.Scan(&v)
		applied[v] = true
	}

	files, err := filepath.Glob("../migrations/*.sql")
	if err != nil {
		log.Fatalf("Failed to list migration files: %v", err)
	}

	sort.Strings(files)

	for _, file := range files {
		version := filepath.Base(file)
		if applied[version] {
			log.Printf("⏩ Skipping already applied migration: %s", version)
			continue
		}

		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read migration %s: %v", version, err)
		}

		sqlStmt := string(sqlBytes)
		if strings.TrimSpace(sqlStmt) == "" {
			log.Printf("⚠️ Empty migration file skipped: %s", version)
			continue
		}

		if _, err := DB.Exec(sqlStmt); err != nil {
			log.Fatalf("❌ Failed to apply migration %s: %v", version, err)
		}

		if _, err := DB.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version); err != nil {
			log.Fatalf("Failed to record migration %s: %v", version, err)
		}

		log.Printf("✅ Applied migration: %s", version)
	}
}
