#!/bin/bash

# Read user input
read -p "Enter the number of rows: " rows

# Initialize counter
counter=1

# Loop to print the pattern
for (( i=1; i<=rows; i++ ))
do
  for (( j=1; j<=i; j++ ))
  do
    echo -n "$counter "
    ((counter++))
  done
  echo ""
done
