package main

import (
	"log"
	"testing"
)

func TestAnswer(t *testing.T) {
	lines, err := readLines("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer := Answer1(lines)
	expected := 15
	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}

func TestAnswer2(t *testing.T) {
	lines, err := readLines("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer := Answer2(lines)
	expected := 12
	if answer != 12 {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
