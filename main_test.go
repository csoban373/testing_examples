package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSomething(t *testing.T) {
// 	expected := "hello world"
// 	output := Something()

// 	if expected != output {
// 		t.Fatal("should equal one another!")
// 	}
// }

func TestSomething(t *testing.T) {
	expected := "hello world"
	output := Something()

	assert.Equal(t, expected, output, "should equal one another!")
}

func TestAddNums(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "example failed test",
			input:    2,
			expected: 5,
		},
		{
			name:     "squaring the input",
			input:    2,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Square(tt.input))
		})
	}
}
