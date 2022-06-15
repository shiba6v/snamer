package snamer

// package is not snamer_test because a test of private method is needed.

import (
	"fmt"
	"testing"
)

func TestPascalStringToCamel(t *testing.T) {
	fmt.Print("TestPascalStringToCamel\n")
	examples := []map[string]string{
		{"input": "ExampleName", "expected": "exampleName"},
		{"input": "A", "expected": "a"},
		{"input": "aA", "expected": "aA"},
		{"input": "AA", "expected": "aA"},
	}
	for _, ex := range examples {
		t.Run(
			"TestPascalStringToCamel: "+ex["input"], func(t *testing.T) {
				t.Parallel()
				s := ex["input"]
				expected := ex["expected"]
				result := pascalStringToCamel(s)
				if result != expected {
					t.Errorf("Error input:%v result: %v, expected:%v", s, result, expected)
				}
			})
	}
}

func TestPascalStringToSnake(t *testing.T) {
	fmt.Print("TestPascalStringToSnake\n")
	examples := []map[string]string{
		{"input": "ExampleName", "expected": "example_name"},
		{"input": "A", "expected": "a"},
		{"input": "aA", "expected": "a_a"},
		{"input": "AA", "expected": "a_a"},
	}
	for _, ex := range examples {
		t.Run(
			"TestPascalStringToCamel: "+ex["input"], func(t *testing.T) {
				t.Parallel()
				s := ex["input"]
				expected := ex["expected"]
				result := pascalStringToSnake(s)
				if result != expected {
					t.Errorf("Error input:%v result: %v, expected:%v", s, result, expected)
				}
			})
	}
}
