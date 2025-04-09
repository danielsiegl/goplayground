# Contract Viewer CLI

A command-line tool for viewing, storing, and managing contracts.

## Features

- Read contract information from JSON files
- Display contract information in the console
- Output contract information as markdown
- Store contracts in SQLite database
- List all contracts in the database
- Delete contracts from the database
- Customizable contract file path

## Usage

```bash
# Display contract information
./goplayground -contract

# Output contract to markdown
./goplayground -output-md

# Store contract in database
./goplayground -store

# List all contracts in database
./goplayground -list

# Delete a contract from database
./goplayground -delete "contract-id"

# Specify a custom contract file
./goplayground -contract -contract-file config/custom-contract.json

# Specify a custom database file
./goplayground -store -db data/custom.db
```

## Available Flags

- `-contract`: Display contract information
- `-output-md`: Output contract to output.md
- `-store`: Store contract in database
- `-list`: List all contracts in database
- `-delete`: Delete contract with the specified ID from database
- `-contract-file`: Path to the contract.json file (default: config/contract.json)
- `-db`: Path to the SQLite database file (default: data/contracts.db)

## Contract Configuration

The program reads contract information from a JSON file. The default contract file is located at `config/contract.json`. You can create custom contract files following the same structure.

Example contract.json:
```json
{
    "id": "CONTRACT-001",
    "title": "Sample Contract",
    "parties": [
        {
            "name": "Alice Johnson",
            "role": "Client",
            "email": "alice@example.com"
        },
        {
            "name": "Bob Wilson",
            "role": "Provider",
            "email": "bob@example.com"
        }
    ],
    "terms": {
        "startDate": "2024-01-01",
        "endDate": "2024-12-31",
        "value": 50000.00,
        "currency": "USD"
    },
    "status": "Active"
}
```

## Database

The program uses SQLite to store contracts. The database file is created at `data/contracts.db` by default. You can specify a custom database file using the `-db` flag.

The database schema includes:
- Contract ID
- Title
- Status
- Parties (stored as JSON)
- Terms (stored as JSON)
- Created timestamp

## Building

Use the build script to compile the program:

```powershell
# Build with default contract file
./build.ps1

# Build with custom contract file
./build.ps1 -ContractFile "config/custom-contract.json"
```

The build script will:
1. Build executables for all platforms
2. Copy the executables to the bin directory
3. Create a config directory inside bin
4. Copy all JSON files from the config directory to bin/config 