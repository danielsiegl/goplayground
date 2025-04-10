# Check if WSL is installed
$wslCheck = Get-Command wsl -ErrorAction SilentlyContinue
if (-not $wslCheck) {
    Write-Error "WSL is not installed. Please install WSL first."
    exit 1
}

# Detect WSL architecture
$wslArch = wsl uname -m
Write-Host "WSL Architecture: $wslArch"

# Determine which binary to use based on architecture
$arch = if ($wslArch -eq "aarch64") { "arm64" } else { "amd64" }
$binaryPath = "bin/goplayground-linux-$arch"
Write-Host "Using binary: $binaryPath"

# Check if the binary exists
if (-not (Test-Path $binaryPath)) {
    Write-Error "Linux binary not found at $binaryPath. Please run build.ps1 first."
    exit 1
}

# Check if contract file exists
$contractPath = "bin/config/contract.json"
if (-not (Test-Path $contractPath)) {
    Write-Error "Contract file not found at $contractPath"
    exit 1
}

# Get the absolute paths
$absoluteBinaryPath = (Get-Item $binaryPath).FullName
$absoluteContractPath = (Get-Item $contractPath).FullName

# Convert Windows paths to WSL paths
$wslBinaryPath = $absoluteBinaryPath -replace '\\', '/' -replace '^C:', '/mnt/c'
$wslContractPath = $absoluteContractPath -replace '\\', '/' -replace '^C:', '/mnt/c'

# Create data directory in WSL if it doesn't exist
Write-Host "Creating data directory in WSL..."
wsl mkdir -p data

# Make the binary executable in WSL
Write-Host "Making binary executable in WSL..."
wsl chmod +x "$wslBinaryPath"

# Run the binary in WSL
Write-Host "Running goplayground in WSL..."
wsl "$wslBinaryPath" -contract-file "$wslContractPath" -store 