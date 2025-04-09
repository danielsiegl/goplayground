package main

import (
	"flag"
	"fmt"
	"os"
)

// defaultContractFile is set at build time using -ldflags
var defaultContractFile string

func main() {
	// Define command-line flags
	showContract := flag.Bool("contract", false, "Show contract information")
	outputMarkdown := flag.Bool("output-md", false, "Output contract to output.md")
	storeContract := flag.Bool("store", false, "Store contract in database")
	listContracts := flag.Bool("list", false, "List all contracts in database")
	deleteContract := flag.String("delete", "", "Delete contract with the specified ID from database")

	// Set default contract file path
	defaultPath := "config/contract.json"
	if defaultContractFile != "" {
		defaultPath = defaultContractFile
	}

	contractFile := flag.String("contract-file", defaultPath, "Path to the contract.json file")
	dbPath := flag.String("db", "data/contracts.db", "Path to the SQLite database file")

	// Parse the flags
	flag.Parse()

	// Initialize database if needed
	var db *DB
	var err error
	if *storeContract || *listContracts || *deleteContract != "" {
		db, err = InitDB(*dbPath)
		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			return
		}
		defer db.Close()
	}

	// Handle delete contract
	if *deleteContract != "" {
		err := db.DeleteContract(*deleteContract)
		if err != nil {
			fmt.Printf("Error deleting contract: %v\n", err)
			return
		}
		fmt.Printf("Contract %s deleted successfully\n", *deleteContract)
		return
	}

	// Handle list contracts
	if *listContracts {
		contracts, err := db.GetAllContracts()
		if err != nil {
			fmt.Printf("Error listing contracts: %v\n", err)
			return
		}

		if len(contracts) == 0 {
			fmt.Println("No contracts found in database")
			return
		}

		fmt.Println("Contracts in database:")
		fmt.Println("ID\tTitle\tStatus")
		fmt.Println("----\t-----\t------")
		for _, contract := range contracts {
			fmt.Printf("%s\t%s\t%s\n", contract.ID, contract.Title, contract.Status)
		}
		return
	}

	// Load contract from file
	contract, err := LoadContract(*contractFile)
	if err != nil {
		fmt.Printf("Error loading contract: %v\n", err)
		return
	}

	// Store contract in database if requested
	if *storeContract {
		err := db.StoreContract(contract)
		if err != nil {
			fmt.Printf("Error storing contract: %v\n", err)
			return
		}
		fmt.Printf("Contract %s stored successfully in database\n", contract.ID)
		return
	}

	// Handle contract display if requested
	if *showContract || *outputMarkdown {
		// Create markdown content
		markdown := fmt.Sprintf("# Contract Information\n\n")
		markdown += fmt.Sprintf("## Basic Information\n")
		markdown += fmt.Sprintf("- **ID:** %s\n", contract.ID)
		markdown += fmt.Sprintf("- **Title:** %s\n", contract.Title)
		markdown += fmt.Sprintf("- **Status:** %s\n\n", contract.Status)

		markdown += fmt.Sprintf("## Parties\n")
		for _, party := range contract.Parties {
			markdown += fmt.Sprintf("### %s (%s)\n", party.Name, party.Role)
			markdown += fmt.Sprintf("- Email: %s\n\n", party.Email)
		}

		markdown += fmt.Sprintf("## Terms\n")
		markdown += fmt.Sprintf("- **Period:** %s to %s\n", contract.Terms.StartDate, contract.Terms.EndDate)
		markdown += fmt.Sprintf("- **Value:** %.2f %s\n", contract.Terms.Value, contract.Terms.Currency)

		// If output to file is requested
		if *outputMarkdown {
			err := os.WriteFile("output.md", []byte(markdown), 0644)
			if err != nil {
				fmt.Printf("Error writing to output.md: %v\n", err)
				return
			}
			fmt.Println("Contract information has been written to output.md")
			return
		}

		// If just displaying to console
		if *showContract {
			fmt.Println(markdown)
			return
		}
	}

	// If no flags are provided, show help
	flag.Usage()
}
