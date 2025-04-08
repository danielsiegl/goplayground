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

	// Set default contract file path
	defaultPath := "config/contract.json"
	if defaultContractFile != "" {
		defaultPath = defaultContractFile
	}

	contractFile := flag.String("contract-file", defaultPath, "Path to the contract.json file")

	// Parse the flags
	flag.Parse()

	// Handle contract display if requested
	if *showContract || *outputMarkdown {
		contract, err := LoadContract(*contractFile)
		if err != nil {
			fmt.Printf("Error loading contract: %v\n", err)
			return
		}

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
