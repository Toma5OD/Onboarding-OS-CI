package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sortMain() {
	// Read lines from a file
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Sort lines
	sort.Strings(lines)

	// Write sorted lines to a new file
	outFile, _ := os.Create("sorted.txt")
	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}
