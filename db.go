package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
)

// DB represents the database connection
type DB struct {
	*sql.DB
}

// InitDB initializes the SQLite database and creates the necessary tables
func InitDB(dbPath string) (*DB, error) {
	// Create the database directory if it doesn't exist
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("error creating database directory: %v", err)
	}

	// Open the database
	db, err := sql.Open(sqlite.DriverName, dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Create the contracts table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS contracts (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL,
		parties_json TEXT NOT NULL,
		terms_json TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating contracts table: %v", err)
	}

	return &DB{db}, nil
}

// StoreContract stores a contract in the database
func (db *DB) StoreContract(contract *Contract) error {
	// Convert parties and terms to JSON
	partiesJSON, err := json.Marshal(contract.Parties)
	if err != nil {
		return fmt.Errorf("error marshaling parties: %v", err)
	}

	termsJSON, err := json.Marshal(contract.Terms)
	if err != nil {
		return fmt.Errorf("error marshaling terms: %v", err)
	}

	// Insert or replace the contract
	query := `
	INSERT OR REPLACE INTO contracts (id, title, status, parties_json, terms_json)
	VALUES (?, ?, ?, ?, ?);`

	_, err = db.Exec(query, contract.ID, contract.Title, contract.Status, string(partiesJSON), string(termsJSON))
	if err != nil {
		return fmt.Errorf("error storing contract: %v", err)
	}

	return nil
}

// GetContract retrieves a contract from the database by ID
func (db *DB) GetContract(id string) (*Contract, error) {
	query := `
	SELECT id, title, status, parties_json, terms_json
	FROM contracts
	WHERE id = ?;`

	var contract Contract
	var partiesJSON, termsJSON string

	err := db.QueryRow(query, id).Scan(&contract.ID, &contract.Title, &contract.Status, &partiesJSON, &termsJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("contract not found: %s", id)
		}
		return nil, fmt.Errorf("error retrieving contract: %v", err)
	}

	// Parse parties JSON
	if err := json.Unmarshal([]byte(partiesJSON), &contract.Parties); err != nil {
		return nil, fmt.Errorf("error unmarshaling parties: %v", err)
	}

	// Parse terms JSON
	if err := json.Unmarshal([]byte(termsJSON), &contract.Terms); err != nil {
		return nil, fmt.Errorf("error unmarshaling terms: %v", err)
	}

	return &contract, nil
}

// GetAllContracts retrieves all contracts from the database
func (db *DB) GetAllContracts() ([]*Contract, error) {
	query := `
	SELECT id, title, status, parties_json, terms_json
	FROM contracts
	ORDER BY created_at DESC;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying contracts: %v", err)
	}
	defer rows.Close()

	var contracts []*Contract
	for rows.Next() {
		var contract Contract
		var partiesJSON, termsJSON string

		err := rows.Scan(&contract.ID, &contract.Title, &contract.Status, &partiesJSON, &termsJSON)
		if err != nil {
			return nil, fmt.Errorf("error scanning contract: %v", err)
		}

		// Parse parties JSON
		if err := json.Unmarshal([]byte(partiesJSON), &contract.Parties); err != nil {
			return nil, fmt.Errorf("error unmarshaling parties: %v", err)
		}

		// Parse terms JSON
		if err := json.Unmarshal([]byte(termsJSON), &contract.Terms); err != nil {
			return nil, fmt.Errorf("error unmarshaling terms: %v", err)
		}

		contracts = append(contracts, &contract)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating contracts: %v", err)
	}

	return contracts, nil
}

// DeleteContract deletes a contract from the database by ID
func (db *DB) DeleteContract(id string) error {
	query := `DELETE FROM contracts WHERE id = ?;`

	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting contract: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("contract not found: %s", id)
	}

	return nil
}
