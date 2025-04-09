package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDB(t *testing.T) {
	// Create a temporary database file for testing
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")

	// Initialize the database
	db, err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Test contract
	testContract := &Contract{
		ID:     "TEST-001",
		Title:  "Test Contract",
		Status: "active",
		Parties: []Party{
			{
				Name:  "Test Party 1",
				Role:  "Client",
				Email: "test1@example.com",
			},
			{
				Name:  "Test Party 2",
				Role:  "Provider",
				Email: "test2@example.com",
			},
		},
		Terms: Terms{
			StartDate: "2024-01-01",
			EndDate:   "2024-12-31",
			Value:     1000.00,
			Currency:  "USD",
		},
	}

	// Test storing a contract
	t.Run("StoreContract", func(t *testing.T) {
		err := db.StoreContract(testContract)
		if err != nil {
			t.Errorf("Failed to store contract: %v", err)
		}
	})

	// Test retrieving a contract
	t.Run("GetContract", func(t *testing.T) {
		contract, err := db.GetContract(testContract.ID)
		if err != nil {
			t.Errorf("Failed to get contract: %v", err)
		}

		// Verify contract fields
		if contract.ID != testContract.ID {
			t.Errorf("Expected ID %s, got %s", testContract.ID, contract.ID)
		}
		if contract.Title != testContract.Title {
			t.Errorf("Expected Title %s, got %s", testContract.Title, contract.Title)
		}
		if contract.Status != testContract.Status {
			t.Errorf("Expected Status %s, got %s", testContract.Status, contract.Status)
		}

		// Verify parties
		if len(contract.Parties) != len(testContract.Parties) {
			t.Errorf("Expected %d parties, got %d", len(testContract.Parties), len(contract.Parties))
		}
		for i, party := range contract.Parties {
			if party.Name != testContract.Parties[i].Name {
				t.Errorf("Expected party name %s, got %s", testContract.Parties[i].Name, party.Name)
			}
			if party.Role != testContract.Parties[i].Role {
				t.Errorf("Expected party role %s, got %s", testContract.Parties[i].Role, party.Role)
			}
			if party.Email != testContract.Parties[i].Email {
				t.Errorf("Expected party email %s, got %s", testContract.Parties[i].Email, party.Email)
			}
		}

		// Verify terms
		if contract.Terms.StartDate != testContract.Terms.StartDate {
			t.Errorf("Expected start date %s, got %s", testContract.Terms.StartDate, contract.Terms.StartDate)
		}
		if contract.Terms.EndDate != testContract.Terms.EndDate {
			t.Errorf("Expected end date %s, got %s", testContract.Terms.EndDate, contract.Terms.EndDate)
		}
		if contract.Terms.Value != testContract.Terms.Value {
			t.Errorf("Expected value %f, got %f", testContract.Terms.Value, contract.Terms.Value)
		}
		if contract.Terms.Currency != testContract.Terms.Currency {
			t.Errorf("Expected currency %s, got %s", testContract.Terms.Currency, contract.Terms.Currency)
		}
	})

	// Test listing all contracts
	t.Run("GetAllContracts", func(t *testing.T) {
		contracts, err := db.GetAllContracts()
		if err != nil {
			t.Errorf("Failed to get all contracts: %v", err)
		}

		if len(contracts) != 1 {
			t.Errorf("Expected 1 contract, got %d", len(contracts))
		}

		contract := contracts[0]
		if contract.ID != testContract.ID {
			t.Errorf("Expected ID %s, got %s", testContract.ID, contract.ID)
		}
	})

	// Test deleting a contract
	t.Run("DeleteContract", func(t *testing.T) {
		err := db.DeleteContract(testContract.ID)
		if err != nil {
			t.Errorf("Failed to delete contract: %v", err)
		}

		// Verify contract is deleted
		_, err = db.GetContract(testContract.ID)
		if err == nil {
			t.Error("Expected error when getting deleted contract, got nil")
		}
	})

	// Test getting non-existent contract
	t.Run("GetNonExistentContract", func(t *testing.T) {
		_, err := db.GetContract("NON-EXISTENT")
		if err == nil {
			t.Error("Expected error when getting non-existent contract, got nil")
		}
	})

	// Test deleting non-existent contract
	t.Run("DeleteNonExistentContract", func(t *testing.T) {
		err := db.DeleteContract("NON-EXISTENT")
		if err == nil {
			t.Error("Expected error when deleting non-existent contract, got nil")
		}
	})
}

func TestInitDB(t *testing.T) {
	// Test with invalid path
	t.Run("InvalidPath", func(t *testing.T) {
		// Use a path with null bytes which is invalid on all operating systems
		_, err := InitDB("invalid\000path/db.sqlite")
		if err == nil {
			t.Error("Expected error with invalid path, got nil")
		}
	})

	// Test with valid path
	t.Run("ValidPath", func(t *testing.T) {
		tmpDir := t.TempDir()
		dbPath := filepath.Join(tmpDir, "test.db")

		db, err := InitDB(dbPath)
		if err != nil {
			t.Errorf("Failed to initialize database: %v", err)
		}
		defer db.Close()

		// Verify database file exists
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			t.Error("Database file was not created")
		}
	})
}
