# Script for installing Node.js and npm

# Create temporary directory
$tempDir = "$env:TEMP\nodejs-install"
New-Item -ItemType Directory -Path $tempDir -Force | Out-Null

# URL for downloading Node.js
$nodeJsUrl = "https://nodejs.org/dist/v20.11.1/node-v20.11.1-x64.msi"
$installerPath = "$tempDir\nodejs-installer.msi"

# Download installer
Write-Host "Downloading Node.js..."
Invoke-WebRequest -Uri $nodeJsUrl -OutFile $installerPath

# Install Node.js
Write-Host "Installing Node.js..."
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i", $installerPath, "/quiet", "/norestart" -Wait

# Check installation
Write-Host "Checking installation..."
$env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path", "User")
try {
    $nodeVersion = node -v
    $npmVersion = npm -v
    Write-Host "Node.js installed: $nodeVersion"
    Write-Host "npm installed: $npmVersion"
} catch {
    Write-Host "Error checking Node.js and npm installation"
    exit 1
}

# Cleanup
Write-Host "Cleaning up temporary files..."
Remove-Item -Path $tempDir -Recurse -Force

Write-Host "Installation completed successfully!" 