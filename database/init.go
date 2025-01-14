package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFile = "events.db"
)

type EventStore struct {
	db *sql.DB
}

// NewEventStore creates a new event store instance
func NewEventStore() (*EventStore, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	store := &EventStore{db: db}
	if err := store.InitTables(); err != nil {
		return nil, err
	}

	return store, nil
}

// Close closes the database connection
func (s *EventStore) Close() error {
	return s.db.Close()
}

// initTables creates required tables if they don't exist
func (s *EventStore) InitTables() error {
	// Model Events table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS model_events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_type TEXT NOT NULL,
			model_id INTEGER,
			uploader TEXT,
			model_name TEXT,
			model_version TEXT,
			model_info TEXT,
			node TEXT,
			block_number INTEGER,
			tx_hash TEXT,
			timestamp INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`); err != nil {
		return fmt.Errorf("failed to create model_events table: %v", err)
	}

	// Workload Events table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS workload_events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_type TEXT NOT NULL,
			session_id INTEGER,
			reporter TEXT,
			worker TEXT,
			epoch_id INTEGER,
			workload INTEGER,
			model_id INTEGER,
			block_number INTEGER,
			tx_hash TEXT,
			timestamp INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`); err != nil {
		return fmt.Errorf("failed to create workload_events table: %v", err)
	}

	// Node Registry Events table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS node_registry_events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_type TEXT NOT NULL,
			wallet TEXT,
			identifier TEXT,
			time INTEGER,
			alias_identifier TEXT,
			block_number INTEGER,
			tx_hash TEXT,
			timestamp INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`); err != nil {
		return fmt.Errorf("failed to create node_registry_events table: %v", err)
	}

	return nil
}
