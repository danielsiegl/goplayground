param(
    [string]$Title,
    [string]$MatrixOS,
    [string]$MatrixShell,
    [string]$MatrixBinary,
    [string]$MatrixArch
)

# Determine OS and Architecture
$osPlatform = [System.Runtime.InteropServices.RuntimeInformation]::OSDescription
$architecture = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture

Write-Output "Running on:"
Write-Output "OS Platform: $osPlatform"
Write-Output "Architecture: $architecture"

# Create job summary
$summary = "$Title`n"
$summary += "OS Platform: $osPlatform`n"
$summary += "Architecture: $architecture`n"
$summary += "Matrix Configuration`n"
$summary += "OS: $MatrixOS`n"
$summary += "Shell: $MatrixShell`n"
$summary += "Binary: $MatrixBinary`n"
$summary += "Architecture: $MatrixArch"

# Write to job summary
$summary | Out-File -FilePath $env:GITHUB_STEP_SUMMARY -Append 