package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLoadContract(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	t.Run("ValidContract", func(t *testing.T) {
		// Create a valid contract JSON file
		contract := Contract{
			ID:     "TEST-001",
			Title:  "Test Contract",
			Status: "active",
			Parties: []Party{
				{
					Name:  "Test Party 1",
					Role:  "Client",
					Email: "test1@example.com",
				},
			},
			Terms: Terms{
				StartDate: "2024-01-01",
				EndDate:   "2024-12-31",
				Value:     1000.00,
				Currency:  "USD",
			},
		}

		contractPath := filepath.Join(tmpDir, "valid-contract.json")
		contractJSON, err := json.MarshalIndent(contract, "", "  ")
		if err != nil {
			t.Fatalf("Failed to marshal contract: %v", err)
		}

		err = os.WriteFile(contractPath, contractJSON, 0644)
		if err != nil {
			t.Fatalf("Failed to write contract file: %v", err)
		}

		// Test loading the contract
		loadedContract, err := LoadContract(contractPath)
		if err != nil {
			t.Errorf("Failed to load valid contract: %v", err)
		}

		// Verify contract fields
		if loadedContract.ID != contract.ID {
			t.Errorf("Expected ID %s, got %s", contract.ID, loadedContract.ID)
		}
		if loadedContract.Title != contract.Title {
			t.Errorf("Expected Title %s, got %s", contract.Title, loadedContract.Title)
		}
		if loadedContract.Status != contract.Status {
			t.Errorf("Expected Status %s, got %s", contract.Status, loadedContract.Status)
		}
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Create an invalid JSON file
		invalidPath := filepath.Join(tmpDir, "invalid.json")
		err := os.WriteFile(invalidPath, []byte("invalid json"), 0644)
		if err != nil {
			t.Fatalf("Failed to write invalid file: %v", err)
		}

		_, err = LoadContract(invalidPath)
		if err == nil {
			t.Error("Expected error loading invalid JSON, got nil")
		}
	})

	t.Run("NonexistentFile", func(t *testing.T) {
		_, err := LoadContract(filepath.Join(tmpDir, "nonexistent.json"))
		if err == nil {
			t.Error("Expected error loading nonexistent file, got nil")
		}
	})

	t.Run("EmptyFile", func(t *testing.T) {
		emptyPath := filepath.Join(tmpDir, "empty.json")
		err := os.WriteFile(emptyPath, []byte{}, 0644)
		if err != nil {
			t.Fatalf("Failed to write empty file: %v", err)
		}

		_, err = LoadContract(emptyPath)
		if err == nil {
			t.Error("Expected error loading empty file, got nil")
		}
	})
}

func TestContractValidation(t *testing.T) {
	t.Run("ValidContract", func(t *testing.T) {
		contract := Contract{
			ID:     "TEST-001",
			Title:  "Test Contract",
			Status: "active",
			Parties: []Party{
				{
					Name:  "Test Party",
					Role:  "Client",
					Email: "test@example.com",
				},
			},
			Terms: Terms{
				StartDate: time.Now().Format("2006-01-02"),
				EndDate:   time.Now().AddDate(1, 0, 0).Format("2006-01-02"),
				Value:     1000.00,
				Currency:  "USD",
			},
		}

		if err := contract.Validate(); err != nil {
			t.Errorf("Expected valid contract to pass validation: %v", err)
		}
	})

	t.Run("MissingID", func(t *testing.T) {
		contract := Contract{
			Title:  "Test Contract",
			Status: "active",
		}

		if err := contract.Validate(); err == nil {
			t.Error("Expected error for missing ID")
		}
	})

	t.Run("InvalidEmail", func(t *testing.T) {
		contract := Contract{
			ID:     "TEST-001",
			Title:  "Test Contract",
			Status: "active",
			Parties: []Party{
				{
					Name:  "Test Party",
					Role:  "Client",
					Email: "invalid-email",
				},
			},
		}

		if err := contract.Validate(); err == nil {
			t.Error("Expected error for invalid email")
		}
	})

	t.Run("InvalidDates", func(t *testing.T) {
		contract := Contract{
			ID:     "TEST-001",
			Title:  "Test Contract",
			Status: "active",
			Terms: Terms{
				StartDate: "2024-12-31",
				EndDate:   "2024-01-01", // End date before start date
				Value:     1000.00,
				Currency:  "USD",
			},
		}

		if err := contract.Validate(); err == nil {
			t.Error("Expected error for invalid dates")
		}
	})

	t.Run("InvalidValue", func(t *testing.T) {
		contract := Contract{
			ID:     "TEST-001",
			Title:  "Test Contract",
			Status: "active",
			Terms: Terms{
				StartDate: "2024-01-01",
				EndDate:   "2024-12-31",
				Value:     -1000.00, // Negative value
				Currency:  "USD",
			},
		}

		if err := contract.Validate(); err == nil {
			t.Error("Expected error for negative value")
		}
	})

	t.Run("InvalidCurrency", func(t *testing.T) {
		contract := Contract{
			ID:     "TEST-001",
			Title:  "Test Contract",
			Status: "active",
			Terms: Terms{
				StartDate: "2024-01-01",
				EndDate:   "2024-12-31",
				Value:     1000.00,
				Currency:  "INVALID", // Invalid currency code
			},
		}

		if err := contract.Validate(); err == nil {
			t.Error("Expected error for invalid currency")
		}
	})
}

func TestContractToMarkdown(t *testing.T) {
	contract := Contract{
		ID:     "TEST-001",
		Title:  "Test Contract",
		Status: "active",
		Parties: []Party{
			{
				Name:  "Test Client",
				Role:  "Client",
				Email: "client@example.com",
			},
			{
				Name:  "Test Provider",
				Role:  "Provider",
				Email: "provider@example.com",
			},
		},
		Terms: Terms{
			StartDate: "2024-01-01",
			EndDate:   "2024-12-31",
			Value:     1000.00,
			Currency:  "USD",
		},
	}

	markdown := contract.ToMarkdown()

	// Test markdown content
	expectedStrings := []string{
		"# Contract Details",
		"## Contract TEST-001",
		"Status: active",
		"### Parties",
		"* Test Client (Client)",
		"* Test Provider (Provider)",
		"### Terms",
		"* Period: 2024-01-01 to 2024-12-31",
		"* Value: 1000.00 USD",
	}

	for _, expected := range expectedStrings {
		if !contains(markdown, expected) {
			t.Errorf("Expected markdown to contain '%s'", expected)
		}
	}
}

// Helper function to check if a string contains another string
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
