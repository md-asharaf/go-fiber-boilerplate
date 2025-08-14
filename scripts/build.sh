#!/bin/bash

# Build script for the Go API

set -e

echo "Building Go API..."

# Clean up previous builds
rm -rf build/

# Create build directory
mkdir -p build/

# Build the application
go build -o build/api ./cmd/main.go

echo "Build completed successfully!"
echo "Binary location: build/api"

# Make it executable
chmod +x build/api

echo "To run the application: ./build/api"
