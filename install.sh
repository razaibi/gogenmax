#!/bin/bash

# Ensure the script stops if there is an error
set -e

# Build the executable
go build -o gmx

# Move the executable to /usr/local/bin
sudo mv gmx /usr/local/bin/

echo "gmx has been installed to /usr/local/bin"