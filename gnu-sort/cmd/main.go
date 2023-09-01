package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// ReadLines reads lines from an io.Reader and returns them as a slice of strings.
func ReadLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// SortLines sorts a slice of strings in ascending order.
func SortLines(lines []string) {
	sort.Strings(lines)
}

// WriteLines writes a slice of strings to an io.Writer.
func WriteLines(w io.Writer, lines []string) error {
	writer := bufio.NewWriter(w)

	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return
	}
	defer file.Close()

	lines, err := ReadLines(file)
	if err != nil {
		log.Printf("Error reading lines: %s", err)
		return
	}

	SortLines(lines)

	outFile, err := os.Create("../sorted-file/sorted.txt")
	if err != nil {
		log.Printf("Error writing to file: %s", err)
		return
	}
	defer outFile.Close()

	err = WriteLines(outFile, lines)
	if err != nil {
		log.Printf("Error writing lines: %s", err)
	}
}
