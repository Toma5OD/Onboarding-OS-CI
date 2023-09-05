package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"Sort fruits", []string{"banana", "apple", "cherry"}, []string{"apple", "banana", "cherry"}},
		{"Sort small words", []string{"cat", "apple", "bat"}, []string{"apple", "bat", "cat"}},
		{"Sort animals", []string{"dog", "cat", "bird"}, []string{"bird", "cat", "dog"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SortLines(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("SortLines = %v; want %v", test.input, test.expected)
			}
		})
	}
}

func BenchmarkSortLines(b *testing.B) {
	lines := []string{"banana", "apple", "cherry", "date", "elderberry", "fig", "grape", "honeydew"}

	// Randomize the lines slice here
	for i := range lines {
		j := rand.Intn(i + 1)
		lines[i], lines[j] = lines[j], lines[i]
	}

	// Reset the timer to exclude setup time
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SortLines(lines)
	}
}
