# Simple Go CLI Program

A simple command-line interface program written in Go that demonstrates basic CLI features including flags and argument handling.

## Features

- Customizable greeting with name
- Configurable number of repetitions
- Option to display greeting in uppercase
- Support for additional arguments

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
```

## Available Flags

- `-name`: Name to greet (default: "World")
- `-count`: Number of times to print the greeting (default: 1)
- `-uppercase`: Print greeting in uppercase (default: false)

## Building

To build the program, simply run:
```bash
go build
```

This will create an executable named `goplayground` in the current directory. 