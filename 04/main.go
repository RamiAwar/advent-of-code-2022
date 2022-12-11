package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type Interval struct {
	Min uint8
	Max uint8
}

// Create interval from string
func NewInterval(s string) *Interval {
	b := strings.IndexByte(s, '-')
	if b == -1 {
		log.Fatal("invalid interval string")
	}

	min, err := ParseInt(s[:b])
	if err != nil {
		log.Fatal(err)
	}

	max, err := ParseInt(s[b+1:])
	if err != nil {
		log.Fatal(err)
	}

	return &Interval{min, max}
}

// Check interval contains other interval
func (i *Interval) Contains(other *Interval) bool {
	return i.Min <= other.Min && i.Max >= other.Max
}

// Check interval overlap
func (i *Interval) Overlaps(other *Interval) bool {
	return (i.Min <= other.Max && i.Max >= other.Min) || (other.Min <= i.Max && other.Max >= i.Min)
}

func Answer1(lines []string) int {
	nContains := 0
	for _, line := range lines {
		i := strings.IndexByte(line, ',')
		if i == -1 {
			continue
		}

		a := NewInterval(line[:i])
		b := NewInterval(line[i+1:])

		if a.Contains(b) || b.Contains(a) {
			nContains++
		}
	}
	return nContains
}

func Answer2(lines []string) int {
	nOverlaps := 0
	for _, line := range lines {
		i := strings.IndexByte(line, ',')
		if i == -1 {
			continue
		}

		a := NewInterval(line[:i])
		b := NewInterval(line[i+1:])

		if a.Overlaps(b) {
			nOverlaps++
		}
	}
	return nOverlaps
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer := Answer1(lines)
	fmt.Println("answer 1:", answer)

	answer2 := Answer2(lines)
	fmt.Println("answer 2:", answer2)
}
