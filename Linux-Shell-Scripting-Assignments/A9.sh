#!/bin/bash

# Check if a filename is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <filename>"
  exit 1
fi

# Validate if file exists
if [ ! -f "$1" ]; then
  echo "File not found!"
  exit 1
fi

# Use sed to remove empty lines
sed -i '/^$/d' "$1"

echo "Empty lines removed from $1."
