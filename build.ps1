function Build-GoApplication {
    param (
        [string]$GOOS,
        [string]$GOARCH = "amd64",
        [string]$CGO_ENABLED = "0"
    )

    # Store original environment variables
    $originalGOOS = $env:GOOS
    $originalGOARCH = $env:GOARCH
    $originalCGO_ENABLED = $env:CGO_ENABLED

    try {
        # Set environment variables for build
        $env:GOOS = $GOOS
        $env:GOARCH = $GOARCH
        $env:CGO_ENABLED = $CGO_ENABLED

        Write-Output "Building for $env:GOOS"
        
        # Determine output filename based on OS
        $outputFile = if ($GOOS -eq "windows") {
            "goplayground-$GOOS-$GOARCH.exe"
        } else {
            "goplayground-$GOOS-$GOARCH"
        }

        # Perform the build
        go build -o $outputFile
        Write-Output "Created: '$outputFile'"
        Write-Output "$GOOS build complete"
    }
    finally {
        # Restore original environment variables
        $env:GOOS = $originalGOOS
        $env:GOARCH = $originalGOARCH
        $env:CGO_ENABLED = $originalCGO_ENABLED
    }
}

# Build for Linux
Build-GoApplication -GOOS "linux"
Build-GoApplication -GOOS "linux" -GOARCH "arm64"

# Build for Windows
Build-GoApplication -GOOS "windows"
Build-GoApplication -GOOS "windows" -GOARCH "arm64"

# Build for macOS on M1,M2,M3,..
Build-GoApplication -GOOS "darwin" -GOARCH "arm64"


