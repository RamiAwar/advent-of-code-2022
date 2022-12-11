package main

import (
	"log"
	"testing"
)

var input, _ = readLines("input.txt")
var testInput, _ = readLines("test_input.txt")

func TestSplitInHalf(t *testing.T) {
	s := "abcdef"
	halves, err := SplitInHalf(s)
	if err != nil {
		log.Fatal(err)
	}
	expected := []string{"abc", "def"}
	if halves[0] != expected[0] || halves[1] != expected[1] {
		t.Errorf("expected %v, got %v", expected, halves)
	}
}

func TestSplitInHalfWithError(t *testing.T) {
	s := "abcdefg"
	_, err := SplitInHalf(s)

	if err == nil {
		t.Errorf("expected odd string length error, got nil")
	}
}

func TestGetPriority(t *testing.T) {
	type Args struct {
		r        rune
		priority int
	}

	args := []Args{
		{r: 'a', priority: 1},
		{r: 'z', priority: 26},
		{r: 'A', priority: 27},
		{r: 'Z', priority: 52},
	}

	for _, arg := range args {
		priority := getPriority(arg.r)
		if priority != arg.priority {
			t.Errorf("expected %d, got %d", arg.priority, priority)
		}
	}
}

func TestAnswer(t *testing.T) {
	answer := Answer1(testInput)
	expected := 157
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}

func TestAnswer2(t *testing.T) {
	answer := Answer2(testInput)
	expected := 70
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
