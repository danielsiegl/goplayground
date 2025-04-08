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
	showUsers := flag.Bool("users", false, "Show users from database")

	// Parse the flags
	flag.Parse()

	// Handle database operations if requested
	if *showUsers {
		db, err := initDB()
		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		users, err := getAllUsers(db)
		if err != nil {
			fmt.Printf("Error getting users: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\nUsers in database:")
		fmt.Println("ID  | Name          | Email")
		fmt.Println("----|---------------|------------------")
		for _, user := range users {
			fmt.Printf("%-4d| %-13s | %s\n", user.ID, user.Name, user.Email)
		}
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
