#!/bin/bash

# Loop for each row
for (( row=1; row<=8; row++ ))
do
  # Loop for each column
  for (( col=1; col<=8; col++ ))
  do
    # Check for alternating colors
    if [ $(( (row + col) % 2 )) -eq 0 ]; then
      echo -e -n "\e[40m  "
    else
      echo -e -n "\e[47m  "
    fi
  done
  # Reset to normal color and move to next line
  echo -e -n "\e[0m\n"
done
