package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define command-line flags
	name := flag.String("name", "World", "Name to greet")
	count := flag.Int("count", 1, "Number of times to print the greeting")
	uppercase := flag.Bool("uppercase", false, "Print greeting in uppercase")
	showContract := flag.Bool("contract", false, "Show contract information")
	outputMarkdown := flag.Bool("output-md", false, "Output contract to output.md")

	// Parse the flags
	flag.Parse()

	// Handle contract display if requested
	if *showContract || *outputMarkdown {
		contract, err := LoadContract()
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

	// Create the greeting message
	greeting := fmt.Sprintf("Hello, %s!", *name)

	// Apply uppercase if requested
	if *uppercase {
		greeting = strings.ToUpper(greeting)
	}

	// Print the greeting the specified number of times
	for i := 0; i < *count; i++ {
		fmt.Println(greeting)
	}

	// Print any remaining arguments
	if flag.NArg() > 0 {
		fmt.Println("\nAdditional arguments:")
		for _, arg := range flag.Args() {
			fmt.Printf("- %s\n", arg)
		}
	}
}
