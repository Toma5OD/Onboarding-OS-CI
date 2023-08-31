package main

import (
	"reflect"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"banana", "apple", "cherry"}, []string{"apple", "banana", "cherry"}},
		{[]string{"cat", "apple", "bat"}, []string{"apple", "bat", "cat"}},
		{[]string{"dog", "cat", "bird"}, []string{"bird", "cat", "dog"}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			SortLines(test.input) // Call your wrapper function
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("SortLines = %v; want %v", test.input, test.expected)
			}
		})
	}
}
