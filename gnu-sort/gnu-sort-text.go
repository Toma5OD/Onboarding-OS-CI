package main

import (
	"bufio"
	"fmt"
	"os"   // For operating system-level functionalities like file handling
	"sort" // For sorting functionalities
)

func main() {
	// Open the file "input.txt" in read-only mode
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Declare a slice to hold the lines from the file
	var lines []string
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Scan each line and append it to the 'lines' slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Sort the 'lines' slice in ascending alphabetical order
	sort.Strings(lines)

	// Create a new file named "sorted.txt" in write-only mode
	outFile, _ := os.Create("sorted.txt")
	// Create a buffered writer for efficient file writing
	writer := bufio.NewWriter(outFile)
	// Write each sorted line to the new file
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	// Flush the writer buffer to make sure all lines are written to the file
	writer.Flush()
}
