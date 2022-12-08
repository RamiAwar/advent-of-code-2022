package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
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

func StringToSet(s string) mapset.Set[rune] {
	return mapset.NewSet([]rune(s)...)
}

func getPriority(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	} else {
		return int(r-'a') + 1
	}
}

func Answer1(lines []string) int {
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		halves, err := SplitInHalf(line)
		if err != nil {
			log.Fatal(err)
		}

		a := StringToSet(halves[0])
		b := StringToSet(halves[1])

		commonItem, ok := a.Intersect(b).Pop()
		if !ok {
			log.Fatal("no intersection found between both halves")
		}

		sum += getPriority(commonItem)
	}
	return sum
}

func processGroup(groups [3]mapset.Set[rune]) int {
	// Find badge item
	badge, ok := groups[0].Intersect(groups[1]).Intersect(groups[2]).Pop()
	if !ok {
		log.Fatalf("No intersection found")
	}

	return getPriority(badge)
}

func Answer2(lines []string) int {
	sum := 0
	var groups [3]mapset.Set[rune]

	for i, line := range lines {
		if line == "" {
			continue
		}

		// Save line as set
		groups[i%3] = StringToSet(line)

		if i != 0 && (i+1)%3 == 0 {
			sum += processGroup(groups)
		}

	}
	return sum
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
