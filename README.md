# Simple Go CLI Program

A simple command-line interface program written in Go that demonstrates basic CLI features including flags, argument handling, and JSON file parsing.

## Features

- Customizable greeting with name
- Configurable number of repetitions
- Option to display greeting in uppercase
- Support for additional arguments
- Contract information display from JSON file

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

# Show contract information
./goplayground -contract
```

## Available Flags

- `-name`: Name to greet (default: "World")
- `-count`: Number of times to print the greeting (default: 1)
- `-uppercase`: Print greeting in uppercase (default: false)
- `-contract`: Show contract information from config/contract.json (default: false)

## Contract Configuration

The program can read contract information from a JSON file located at `config/contract.json`. The contract file should follow this structure:

```json
{
  "id": "CONTRACT-001",
  "title": "Sample Contract",
  "parties": [
    {
      "name": "John Doe",
      "role": "buyer",
      "email": "john@example.com"
    },
    {
      "name": "Jane Smith",
      "role": "seller",
      "email": "jane@example.com"
    }
  ],
  "terms": {
    "startDate": "2023-01-01",
    "endDate": "2023-12-31",
    "value": 50000.00,
    "currency": "USD"
  },
  "status": "active"
}
```

## Building

To build the program, simply run:
```bash
go build
```

This will create an executable named `goplayground` in the current directory. 