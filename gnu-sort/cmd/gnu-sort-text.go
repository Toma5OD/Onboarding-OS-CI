package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// ReadLines reads lines from a file and returns them as a slice of strings.
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// SortLines sorts a slice of strings in ascending order.
func SortLines(lines []string) {
	sort.Strings(lines)
}

// WriteLines writes a slice of strings to a file.
func WriteLines(filename string, lines []string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
		return err
	}
	writer := bufio.NewWriter(outFile)

	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			log.Fatalf("Error writing line to file: %s", err)
		}
	}
	if err := writer.Flush(); err != nil {
		log.Fatalf("Failed to flush writer: %s", err)
	}

	return nil
}

func main() {
	lines, err := ReadLines("../input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
		return
	}

	SortLines(lines)

	err = WriteLines("sorted.txt", lines)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}
