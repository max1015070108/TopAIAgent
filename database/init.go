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
	// models  table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS models (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    model_id INTEGER NOT NULL UNIQUE,
		    model_name TEXT NOT NULL,
		    model_version TEXT NOT NULL,
		    uploader ADDRESS NOT NULL,
		    model_info TEXT,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL
		)
	`); err != nil {
		return fmt.Errorf("failed to create model_events table: %v", err)
	}

	//model_deployments
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS model_deployments (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    model_id INTEGER NOT NULL,
		    node ADDRESS NOT NULL,
		    status INTEGER NOT NULL,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL,
		    UNIQUE(model_id, node)
		)
	`); err != nil {
		return fmt.Errorf("failed to create model_events table: %v", err)
	}

	// Workload Events table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS workloads (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    session_id INTEGER NOT NULL,
		    epoch_id INTEGER NOT NULL,
		    reporter ADDRESS NOT NULL,
		    worker ADDRESS NOT NULL,
		    model_id INTEGER NOT NULL,
		    workload INTEGER NOT NULL,
		    created_at INTEGER NOT NULL,
		    UNIQUE(session_id, epoch_id, worker)
		)
	`); err != nil {
		return fmt.Errorf("failed to create workload_events table: %v", err)
	}

	// Node Registry Events table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS nodes (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    identifier ADDRESS NOT NULL UNIQUE,
		    wallet ADDRESS NOT NULL,
		    alias_identifier TEXT NOT NULL UNIQUE,
		    registration_time INTEGER NOT NULL,
		    active BOOLEAN NOT NULL DEFAULT FALSE,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL
		);
	`); err != nil {
		return fmt.Errorf("failed to create node_registry_events table: %v", err)
	}

	// Nodes Registry Node table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS node_gpus (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    node_identifier ADDRESS NOT NULL,
		    gpu_type TEXT NOT NULL,
		    total_num INTEGER NOT NULL,
		    used_num INTEGER NOT NULL DEFAULT 0,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL,
		    UNIQUE(node_identifier, gpu_type)
		)
	`); err != nil {
		return fmt.Errorf("failed to create nodes_registry_node table: %v", err)
	}

	// Nodes Registry GPU table for GPU arrays
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS node_relationships (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    provider_identifier ADDRESS NOT NULL,
		    server_identifier ADDRESS NOT NULL,
		    status INTEGER NOT NULL,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL,
		    UNIQUE(provider_identifier, server_identifier)
		)
	`); err != nil {
		return fmt.Errorf("failed to create nodes_registry_gpu table: %v", err)
	}

	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS validations (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    round_id INTEGER NOT NULL,
		    candidate ADDRESS NOT NULL,
		    expected_completion_time INTEGER NOT NULL,
		    status INTEGER NOT NULL,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL,
		    UNIQUE(round_id, candidate)
		)
	`); err != nil {
		return fmt.Errorf("failed to create nodes_registry_gpu table: %v", err)
	}

	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS stakes (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    user ADDRESS NOT NULL,
		    contract_address ADDRESS NOT NULL,
		    amount TEXT NOT NULL,
		    unlock_time INTEGER,
		    status INTEGER NOT NULL,
		    created_at INTEGER NOT NULL,
		    updated_at INTEGER NOT NULL,
		    UNIQUE(user, contract_address)
		)
	`); err != nil {
		return fmt.Errorf("failed to create nodes_registry_gpu table: %v", err)
	}

	// Block Info table
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS block_info (
			id INTEGER PRIMARY KEY CHECK (id = 1),
		    block_hash TEXT NOT NULL,
		    timestamp INTEGER NOT NULL,
		    number INTEGER NOT NULL,
		    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
				`); err != nil {
		return fmt.Errorf("failed to create block_info table: %v", err)
	}

	return nil
}
