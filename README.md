# Simple Go CLI Program

A simple command-line interface program written in Go that demonstrates basic CLI features including flags, argument handling, and JSON file parsing.

## Features

- Customizable greeting with name
- Configurable number of repetitions
- Option to display greeting in uppercase
- Support for additional arguments
- Contract information display from JSON file
- Contract information export to Markdown
- Customizable contract file path

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

# Output contract to markdown file
./goplayground -output-md

# Use a custom contract file
./goplayground -contract -contract-file config/custom-contract.json
```

## Available Flags

- `-name`: Name to greet (default: "World")
- `-count`: Number of times to print the greeting (default: 1)
- `-uppercase`: Print greeting in uppercase (default: false)
- `-contract`: Show contract information from config/contract.json (default: false)
- `-output-md`: Output contract information to output.md (default: false)
- `-contract-file`: Path to the contract.json file (default: "config/contract.json")

## Contract Configuration

The program can read contract information from a JSON file. By default, it looks for `config/contract.json`, but you can specify a different file using the `-contract-file` flag.

### Default Contract File

The default contract file should follow this structure:

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

### Custom Contract Files

You can create custom contract files with different data and specify them at runtime:

```bash
./goplayground -contract -contract-file config/custom-contract.json
```

You can also set a default contract file at build time:

```bash
go build -ldflags "-X main.defaultContractFile=config/custom-contract.json"
```

## Markdown Output

When using the `-output-md` flag, the program will generate an `output.md` file with the contract information formatted in Markdown. The output will include:

- Basic contract information (ID, Title, Status)
- Parties involved (with names, roles, and emails)
- Contract terms (period and value)

The markdown file can be viewed in any markdown viewer or converted to other formats using markdown tools.

## Building

To build the program, simply run:
```bash
go build
```

This will create an executable named `goplayground` in the current directory.

### Cross-Platform Building

The repository includes a PowerShell script (`build.ps1`) that builds the application for multiple platforms:

```powershell
./build.ps1
```

This will create executables for:
- Linux (amd64 and arm64)
- Windows (amd64 and arm64)
- macOS (arm64 for M1/M2/M3)

The built files will be placed in the `bin` directory.

### Building with Custom Contract File

You can also build the application with a custom default contract file:

```powershell
./build.ps1 -ContractFile "config/custom-contract.json"
``` 