package main

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Party represents a party involved in the contract
type Party struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

// Terms represents the contract terms
type Terms struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Value     float64 `json:"value"`
	Currency  string  `json:"currency"`
}

// Contract represents the main contract structure
type Contract struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Parties []Party `json:"parties"`
	Terms   Terms   `json:"terms"`
	Status  string  `json:"status"`
}

// LoadContract reads the contract.json file from the specified path and returns a Contract object
func LoadContract(filePath string) (*Contract, error) {
	// If no file path is provided, use the default path
	if filePath == "" {
		// Get the current working directory
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("error getting working directory: %v", err)
		}

		// Construct the path to the contract.json file
		filePath = filepath.Join(wd, "config", "contract.json")
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading contract file: %v", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("contract file is empty")
	}

	// Parse the JSON into a Contract struct
	var contract Contract
	if err := json.Unmarshal(data, &contract); err != nil {
		return nil, fmt.Errorf("error parsing contract JSON: %v", err)
	}

	if err := contract.Validate(); err != nil {
		return nil, fmt.Errorf("contract validation failed: %v", err)
	}

	return &contract, nil
}

// Validate checks if the contract is valid
func (c *Contract) Validate() error {
	// Check required fields
	if c.ID == "" {
		return fmt.Errorf("contract ID is required")
	}
	if c.Title == "" {
		return fmt.Errorf("contract title is required")
	}
	if c.Status == "" {
		return fmt.Errorf("contract status is required")
	}

	// Validate parties
	if len(c.Parties) == 0 {
		return fmt.Errorf("at least one party is required")
	}
	for _, party := range c.Parties {
		if party.Name == "" {
			return fmt.Errorf("party name is required")
		}
		if party.Role == "" {
			return fmt.Errorf("party role is required")
		}
		if party.Email != "" {
			if _, err := mail.ParseAddress(party.Email); err != nil {
				return fmt.Errorf("invalid email address for party %s: %v", party.Name, err)
			}
		}
	}

	// Validate terms
	if c.Terms.StartDate != "" && c.Terms.EndDate != "" {
		startDate, err := time.Parse("2006-01-02", c.Terms.StartDate)
		if err != nil {
			return fmt.Errorf("invalid start date format: %v", err)
		}
		endDate, err := time.Parse("2006-01-02", c.Terms.EndDate)
		if err != nil {
			return fmt.Errorf("invalid end date format: %v", err)
		}
		if endDate.Before(startDate) {
			return fmt.Errorf("end date cannot be before start date")
		}
	}

	if c.Terms.Value < 0 {
		return fmt.Errorf("contract value cannot be negative")
	}

	// Validate currency (simple check for 3-letter code)
	if c.Terms.Currency != "" && !isValidCurrency(c.Terms.Currency) {
		return fmt.Errorf("invalid currency code: %s", c.Terms.Currency)
	}

	return nil
}

// isValidCurrency checks if the currency code is valid (simple 3-letter check)
func isValidCurrency(currency string) bool {
	// This is a simplified check. In a real application, you might want to
	// check against a list of valid ISO 4217 currency codes
	return len(currency) == 3 && currency == strings.ToUpper(currency)
}

// ToMarkdown converts the contract to markdown format
func (c *Contract) ToMarkdown() string {
	var sb strings.Builder

	sb.WriteString("# Contract Details\n\n")
	sb.WriteString(fmt.Sprintf("## Contract %s\n", c.ID))
	if c.Title != "" {
		sb.WriteString(fmt.Sprintf("Title: %s\n", c.Title))
	}
	sb.WriteString(fmt.Sprintf("Status: %s\n\n", c.Status))

	if len(c.Parties) > 0 {
		sb.WriteString("### Parties\n")
		for _, party := range c.Parties {
			sb.WriteString(fmt.Sprintf("* %s (%s)\n", party.Name, party.Role))
			if party.Email != "" {
				sb.WriteString(fmt.Sprintf("  - Email: %s\n", party.Email))
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("### Terms\n")
	if c.Terms.StartDate != "" && c.Terms.EndDate != "" {
		sb.WriteString(fmt.Sprintf("* Period: %s to %s\n", c.Terms.StartDate, c.Terms.EndDate))
	}
	if c.Terms.Value > 0 {
		sb.WriteString(fmt.Sprintf("* Value: %.2f %s\n", c.Terms.Value, c.Terms.Currency))
	}

	return sb.String()
}
