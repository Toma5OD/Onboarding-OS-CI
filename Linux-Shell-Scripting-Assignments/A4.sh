#!/bin/bash

# Show menu to user
echo "Select an option:"
echo "1) SSH"
echo "2) SCP"
read -p "Enter your choice: " choice

# Read username and IP address
read -p "Enter IP address: " ip_address
read -p "Enter username: " username

# Function for SCP
scp_function() {
  read -p "Direction (1: Remote to Local, 2: Local to Remote): " direction
  read -p "Enter source file location: " src_file

  if [[ -z "$src_file" ]]; then
    echo "Source file location cannot be empty."
    exit 1
  fi

  read -p "Enter destination location (Optional): " dest_location
  dest_file_name=$(basename "$src_file")

  case "$direction" in
    1)
      scp "${username}@${ip_address}:${src_file}" "${dest_location:-./}${dest_file_name}"
      ;;
    2)
      scp "$src_file" "${username}@${ip_address}:${dest_location:-~}/${dest_file_name}"
      ;;
    *)
      echo "Invalid direction."
      ;;
  esac
}

# Perform actions based on choice
case "$choice" in
  1)
    ssh "${username}@${ip_address}"
    ;;
  2)
    scp_function
    ;;
  *)
    echo "Invalid choice."
    ;;
esac
