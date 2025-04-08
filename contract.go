package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

// LoadContract reads the contract.json file from the config directory and returns a Contract object
func LoadContract() (*Contract, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting working directory: %v", err)
	}

	// Construct the path to the contract.json file
	contractPath := filepath.Join(wd, "config", "contract.json")

	// Read the file
	data, err := os.ReadFile(contractPath)
	if err != nil {
		return nil, fmt.Errorf("error reading contract file: %v", err)
	}

	// Parse the JSON into a Contract struct
	var contract Contract
	if err := json.Unmarshal(data, &contract); err != nil {
		return nil, fmt.Errorf("error parsing contract JSON: %v", err)
	}

	return &contract, nil
}
