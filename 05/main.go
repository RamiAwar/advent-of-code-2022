package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gammazero/deque"
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

type Stack struct {
	deque.Deque[rune]
}
type Stacks []Stack

// Print stack
func (s *Stack) String() string {
	stackPrint := make([]rune, 4*(s.Len()+1))
	for i := 0; i < s.Len(); i++ {
		stackPrint = append(stackPrint, '[', s.At(i), ']', ' ')
	}
	return string(stackPrint)
}

// Print stacks
func (s *Stacks) Print() {
	for _, stack := range *s {
		fmt.Println(stack.String())
	}
}

type Move struct {
	From     int
	To       int
	Quantity int
}

// Create move from string
func NewMoveFromString(s string) *Move {
	m := &Move{}
	fmt.Sscanf(s, "move %d from %d to %d", &m.Quantity, &m.From, &m.To)
	return m
}

// Print move
func (m *Move) Print() {
	fmt.Printf("move %d from %d to %d\n", m.Quantity, m.From, m.To)
}

type Moves []*Move

// Print moves
func (m *Moves) Print() {
	for _, move := range *m {
		move.Print()
	}
}

func ProcessInput(lines []string) (*Stacks, *Moves) {
	index := 0
	line := []rune(lines[index])
	stacks := make(Stacks, 10)
	for len(line) != 0 {
		// Fill out list of reversed stacks top to bottom
		maxPosition := len(line) - 1
		runePosition := 0

		for 4*runePosition+1 < maxPosition {
			c := line[4*runePosition+1]
			if c != ' ' {
				stacks[runePosition].PushBack(c)
			}

			runePosition++
		}

		index++
		line = []rune(lines[index])
	}

	// Discard last line (index only)
	for i := range stacks {
		if stacks[i].Len() > 0 {
			stacks[i].PopBack()
		}
	}

	// Begin processing moves
	index++
	moves := make(Moves, 0)
	for index < len(lines) {
		move := NewMoveFromString(lines[index])
		moves = append(moves, move)
		index++
	}

	return &stacks, &moves
}

// Get top of list of stacks as string
func GetTopCrates(stacks *Stacks) string {
	top := make([]rune, 0)
	for i := range *stacks {
		if (*stacks)[i].Len() > 0 {
			top = append(top, (*stacks)[i].PopFront())
		}
	}
	return string(top)
}

func Answer1(stacks *Stacks, moves *Moves) string {
	// Execute moves
	for _, move := range *moves {
		for i := 0; i < move.Quantity; i++ {
			c := (*stacks)[move.From-1].PopFront()
			(*stacks)[move.To-1].PushFront(c)
		}
	}

	return GetTopCrates(stacks)
}

func Answer2(stacks *Stacks, moves *Moves) string {
	// Execute moves while preserving order
	for _, move := range *moves {
		var temp Stack
		for i := 0; i < move.Quantity; i++ {
			c := (*stacks)[move.From-1].PopFront()
			temp.PushFront(c)
		}

		for i := 0; i < move.Quantity; i++ {
			c := temp.PopFront()
			(*stacks)[move.To-1].PushFront(c)
		}
	}

	return GetTopCrates(stacks)
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Process input
	stacks, moves := ProcessInput(lines)

	answer := Answer1(stacks, moves)
	fmt.Println("answer 1:", answer)

	// Process input again (clean)
	stacks, moves = ProcessInput(lines)

	answer2 := Answer2(stacks, moves)
	fmt.Println("answer 2:", answer2)
}
