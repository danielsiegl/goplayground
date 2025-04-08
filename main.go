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

	// Parse the flags
	flag.Parse()

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
