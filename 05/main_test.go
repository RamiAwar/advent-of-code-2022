package main

import (
	"testing"
)

var input, _ = readLines("input.txt")
var testInput, _ = readLines("test_input.txt")

// Process input
var stacks, moves = ProcessInput(testInput)
var stacksTest, movesTest = ProcessInput(testInput)

func TestAnswer(t *testing.T) {
	answer := Answer1(stacksTest, movesTest)
	expected := "CMZ"
	if answer != expected {
		t.Errorf("expected %s, got %s", expected, answer)
	}
}

func TestAnswer2(t *testing.T) {
	answer := Answer2(stacksTest, movesTest)
	expected := "MCD"
	if answer != expected {
		t.Errorf("expected %s, got %s", expected, answer)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var stacks, moves = ProcessInput(testInput)
		Answer1(stacks, moves)
	}
}

func BenchmarkAnswer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var stacks, moves = ProcessInput(testInput)
		Answer2(stacks, moves)
	}
}
