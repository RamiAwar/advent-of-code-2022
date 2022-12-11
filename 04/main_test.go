package main

import (
	"testing"
)

var input, _ = readLines("input.txt")
var testInput, _ = readLines("test_input.txt")

func TestAnswer(t *testing.T) {
	answer := Answer1(testInput)
	expected := 2
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}

func TestAnswer2(t *testing.T) {
	answer := Answer2(testInput)
	expected := 4
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
