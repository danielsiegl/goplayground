package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Define command-line flags
	name := flag.String("name", "World", "Name to greet")
	count := flag.Int("count", 1, "Number of times to print the greeting")
	uppercase := flag.Bool("uppercase", false, "Print greeting in uppercase")
	showContract := flag.Bool("contract", false, "Show contract information")

	// Parse the flags
	flag.Parse()

	// Handle contract display if requested
	if *showContract {
		contract, err := LoadContract()
		if err != nil {
			fmt.Printf("Error loading contract: %v\n", err)
			return
		}

		fmt.Println("\nContract Information:")
		fmt.Printf("ID: %s\n", contract.ID)
		fmt.Printf("Title: %s\n", contract.Title)
		fmt.Printf("Status: %s\n", contract.Status)

		fmt.Println("\nParties:")
		for _, party := range contract.Parties {
			fmt.Printf("- %s (%s): %s\n", party.Name, party.Role, party.Email)
		}

		fmt.Println("\nTerms:")
		fmt.Printf("Period: %s to %s\n", contract.Terms.StartDate, contract.Terms.EndDate)
		fmt.Printf("Value: %.2f %s\n", contract.Terms.Value, contract.Terms.Currency)

		return
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
