package main

import (
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
