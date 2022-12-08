package main

import (
	"errors"
)

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func SplitInHalf(s string) ([]string, error) {
	if len(s)%2 != 0 {
		return []string{}, errors.New("string could not be split in half, odd length")
	}
	return Chunks(s, len(s)/2), nil
}
