package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
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

func Answer1(lines []string) int {
	var current_calories uint64 = 0
	var max_calories uint64 = 0

	for _, depth := range lines {
		if len(depth) == 0 {
			if current_calories > max_calories {
				max_calories = current_calories
			}
			current_calories = 0
		}

		if i, err := strconv.ParseUint(depth, 10, 32); err == nil {
			current_calories += i
		}
	}

	// Deal with edge case if no space at end
	if current_calories > max_calories {
		max_calories = current_calories
	}

	return int(max_calories)
}

func Answer2(lines []string) int {
	var current_calories uint64 = 0

	eq := make(ElfQueue, 0)

	for _, depth := range lines {
		if len(depth) == 0 {
			elf := Elf{calories: int(current_calories)}
			eq.Push(&elf)
			heap.Push(&eq, &elf)
			current_calories = 0
		}

		if i, err := strconv.ParseUint(depth, 10, 32); err == nil {
			current_calories += i
		}
	}

	// Deal with edge case if no space at end
	if current_calories > 0 {
		elf := Elf{calories: int(current_calories)}
		heap.Push(&eq, &elf)
		current_calories = 0
	}

	sum := 0
	for i := 0; i < 3; i++ {
		sum += heap.Pop(&eq).(*Elf).calories
	}

	return sum
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top elf calories: ", Answer1(lines))
	fmt.Println("Total calories by top 3 elves: ", Answer2(lines))
}
