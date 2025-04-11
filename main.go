package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// defaultContractFile is set at build time using -ldflags
var defaultContractFile string

var (
	showContract   = flag.Bool("contract", false, "Show contract information")
	outputMarkdown = flag.Bool("output-md", false, "Output contract information to output.md")
	contractFile   = flag.String("contract-file", "config/contract.json", "Path to the contract.json file")
	storeContract  = flag.Bool("store", false, "Store contract in database")
	listContracts  = flag.Bool("list", false, "List all contracts in database")
	deleteContract = flag.String("delete", "", "Delete contract with the specified ID from database")
	dbPath         = flag.String("db", "data/contracts.db", "Path to the SQLite database file")
)

func main() {
	flag.Parse()

	// Check if any flags are provided
	if flag.NFlag() == 0 {
		fmt.Println("No flags provided. Available options:")
		flag.Usage()
		return
	}

	// Initialize database if any database-related flags are used
	var db *DB
	if *storeContract || *listContracts || *deleteContract != "" {
		var err error
		db, err = InitDB(*dbPath)
		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			return
		}
		defer db.Close()
	}

	// Handle delete operation
	if *deleteContract != "" {
		err := db.DeleteContract(*deleteContract)
		if err != nil {
			fmt.Printf("Error deleting contract: %v\n", err)
			return
		}
		fmt.Printf("Contract %s deleted successfully from database\n", *deleteContract)
		return
	}

	// Handle list operation
	if *listContracts {
		contracts, err := db.GetAllContracts()
		if err != nil {
			fmt.Printf("Error listing contracts: %v\n", err)
			return
		}

		fmt.Println("Contracts in database:")
		fmt.Printf("%-15s %-20s %-10s\n", "ID", "Title", "Status")
		fmt.Println(strings.Repeat("-", 45))
		for _, contract := range contracts {
			fmt.Printf("%-15s %-20s %-10s\n", contract.ID, contract.Title, contract.Status)
		}
		return
	}

	// For store operation, we need a contract file
	if *storeContract {
		// Load the contract from the specified file
		contract, err := LoadContract(*contractFile)
		if err != nil {
			fmt.Printf("Error loading contract from %s: %v\n", *contractFile, err)
			fmt.Println("Please ensure the contract file exists and is valid JSON")
			return
		}

		err = db.StoreContract(contract)
		if err != nil {
			fmt.Printf("Error storing contract: %v\n", err)
			return
		}
		fmt.Printf("Contract %s stored successfully in database\n", contract.ID)
		return
	}

	// Handle contract display if requested
	if *showContract || *outputMarkdown {
		// Load the contract from the specified file
		contract, err := LoadContract(*contractFile)
		if err != nil {
			fmt.Printf("Error loading contract from %s: %v\n", *contractFile, err)
			fmt.Println("Please ensure the contract file exists and is valid JSON")
			return
		}

		markdown := contract.ToMarkdown()

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
}
