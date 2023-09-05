package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

var tempDir string
var tempFile *os.File

func TestMain(m *testing.M) {
	// Setup
	fmt.Println("Setting up tests...")
	if !setup() {
		fmt.Println("Setup failed.")
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	// Teardown
	fmt.Println("Tearing down tests...")
	teardown()

	// Exit
	os.Exit(code)
}

func setup() bool {
	var err error
	tempDir, err = os.MkdirTemp("", "e2e_test")
	if err != nil {
		fmt.Println("Error creating temp directory:", err)
		return false
	}
	tempFile, err = os.Create(fmt.Sprintf("%s/sorted-e2e-test.txt", tempDir))
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return false
	}
	return true
}

func teardown() {
	if err := os.RemoveAll(tempDir); err != nil {
		fmt.Println("Error removing temp directory:", err)
	}
}

func TestEndToEnd(t *testing.T) {
	inputContent := "banana\napple\ncherry"
	inputFile, err := os.CreateTemp("", "input")
	if err != nil {
		t.Fatalf("Could not create temp file: %s", err)
	}
	defer os.Remove(inputFile.Name())
	_, err = inputFile.WriteString(inputContent)
	if err != nil {
		t.Fatalf("Could not write to temp file: %s", err)
	}
	inputFile.Close()

	// Read from the temporary input file
	inputFile, err = os.Open(inputFile.Name())
	if err != nil {
		t.Fatalf("Could not reopen temp input file: %s", err)
	}
	defer inputFile.Close()

	lines, err := ReadLines(inputFile)
	if err != nil {
		t.Fatalf("Error reading lines: %s", err)
	}

	// Sort the lines
	SortLines(lines)

	// Write to the temporary output file
	err = WriteLines(tempFile, lines)
	if err != nil {
		t.Fatalf("Error writing sorted lines: %s", err)
	}

	// Close the temporary output file
	tempFile.Close()

	outputFile, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not open output file: %s", err)
	}
	defer os.Remove(outputFile.Name())

	content, err := io.ReadAll(outputFile)
	if err != nil {
		t.Fatalf("Error reading output file: %s", err)
	}
	lines = strings.Split(string(content), "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sorted := true
	for i := 1; i < len(lines); i++ {
		if lines[i] < lines[i-1] {
			sorted = false
			break
		}
	}

	if !sorted {
		t.Errorf("Lines are not sorted: %v", lines)
	}

	outputFile.Close()
}
