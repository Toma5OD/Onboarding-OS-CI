#!/bin/bash

# Ask user to enter two numbers
read -p "Enter the first number: " num1
read -p "Enter the second number: " num2

# Perform addition using bc and piping
sum=$(echo "$num1 + $num2" | bc)

# Display the result
echo "The sum is: $sum"
