package main

import (
	"testing"
)

var input, _ = readLines("input.txt")
var testInput, _ = readLines("test_input.txt")

func TestAnswer(t *testing.T) {
	answer := Answer1(testInput)
	expected := 15
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}

func TestAnswer2(t *testing.T) {
	answer := Answer2(testInput)
	expected := 12
	if answer != 12 {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer1(input)
	}
}

func BenchmarkAnswer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer2(input)
	}
}
