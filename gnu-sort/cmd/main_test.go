package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Prepare a test file
	testFile := "test_input.txt"
	lines := []string{"apple", "banana", "cherry"}
	file, _ := os.Create(testFile)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	file.Close()

	// Test ReadLines function
	readLines, err := ReadLines(testFile)
	if err != nil {
		t.Errorf("Failed to read lines: %s", err)
	}

	if !reflect.DeepEqual(readLines, lines) {
		t.Errorf("ReadLines = %v; want %v", readLines, lines)
	}

	// Cleanup
	os.Remove(testFile)
}

func TestSortLines(t *testing.T) {
	lines := []string{"banana", "apple", "cherry"}
	expected := []string{"apple", "banana", "cherry"}
	SortLines(lines)

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("SortLines = %v; want %v", lines, expected)
	}
}

func TestWriteLines(t *testing.T) {
	// Prepare lines and test file
	lines := []string{"apple", "banana", "cherry"}
	testFile := "test_output.txt"

	// Test WriteLines function
	err := WriteLines(testFile, lines)
	if err != nil {
		t.Errorf("Failed to write lines: %s", err)
	}

	// Verify if the lines are written correctly
	readLines, err := ReadLines(testFile)
	if err != nil {
		t.Errorf("Failed to read lines: %s", err)
	}

	if !reflect.DeepEqual(readLines, lines) {
		t.Errorf("WriteLines = %v; want %v", readLines, lines)
	}

	// Cleanup
	os.Remove(testFile)
}
