# Simple Go CLI Program

A simple command-line interface program written in Go that demonstrates basic CLI features including flags, argument handling, and SQLite database operations.

## Features

- Customizable greeting with name
- Configurable number of repetitions
- Option to display greeting in uppercase
- Support for additional arguments
- SQLite database integration with sample user data

## Prerequisites

- Go 1.21 or later
- GCC (required for SQLite driver compilation)

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

## Usage

Build the program:
```bash
go build
```

Run the program with various options:

```bash
# Basic usage
./goplayground

# Specify a name
./goplayground -name Alice

# Print greeting multiple times
./goplayground -count 3

# Print in uppercase
./goplayground -uppercase

# Combine options
./goplayground -name Bob -count 2 -uppercase

# Add additional arguments
./goplayground -name Alice extra arg1 arg2

# Show users from database
./goplayground -users
```

## Available Flags

- `-name`: Name to greet (default: "World")
- `-count`: Number of times to print the greeting (default: 1)
- `-uppercase`: Print greeting in uppercase (default: false)
- `-users`: Show users from the SQLite database (default: false)

## Database

The program includes a SQLite database (`users.db`) that is automatically created when you first run the program with the `-users` flag. The database contains a sample table of users with the following fields:
- id (INTEGER, PRIMARY KEY)
- name (TEXT)
- email (TEXT, UNIQUE)

Sample data is automatically inserted when the database is first created.

## Building

To build the program, simply run:
```bash
go build
```

This will create an executable named `goplayground` in the current directory. 