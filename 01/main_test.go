package main

import "testing"

var input, _ = readLines("input.txt")

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
