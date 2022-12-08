package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Move uint8

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

type Result uint8

const (
	Win  Result = 10
	Lose Result = 11
	Draw Result = 12
)

func CalculateScore(opponentMove Move, yourMove Move) uint8 {
	if opponentMove == Rock && yourMove == Scissors {
		return 0
	} else if opponentMove == Scissors && yourMove == Paper {
		return 0
	} else if opponentMove == Paper && yourMove == Scissors {
		return 6
	} else if opponentMove == Paper && yourMove == Rock {
		return 0
	} else if opponentMove == Rock && yourMove == Paper {
		return 6
	} else if opponentMove == Scissors && yourMove == Rock {
		return 6
	}

	return 3 // draw otherwise
}

func CalculateMove(opponentMove Move, result Result) Move {
	if result == Win {
		if opponentMove == Rock {
			return Paper
		} else if opponentMove == Paper {
			return Scissors
		} else if opponentMove == Scissors {
			return Rock
		}
	} else if result == Lose {
		if opponentMove == Rock {
			return Scissors
		} else if opponentMove == Scissors {
			return Paper
		} else if opponentMove == Paper {
			return Rock
		}
	}

	// Draw
	return opponentMove
}

func ParseMove(s rune) (Move, error) {
	if s == 'A' || s == 'X' {
		return Rock, nil
	} else if s == 'B' || s == 'Y' {
		return Paper, nil
	} else if s == 'C' || s == 'Z' {
		return Scissors, nil
	}

	return Move(10), fmt.Errorf("'%q' is not a valid move", s)
}

func ParseResult(s rune) (Result, error) {
	if s == 'X' {
		return Lose, nil
	} else if s == 'Y' {
		return Draw, nil
	} else if s == 'Z' {
		return Win, nil
	}
	return Result(0), fmt.Errorf("'%q' is ot a valid result", s)
}

func Answer1(lines []string) int {
	totalScore := 0

	for _, line := range lines {
		var opponentMove, yourMove Move
		if i, err := ParseMove(rune(line[0])); err == nil {
			opponentMove = i
		}

		if i, err := ParseMove(rune(line[2])); err == nil {
			yourMove = i
		}

		totalScore += int(yourMove) + int(CalculateScore(opponentMove, yourMove))
	}

	return totalScore
}

func Answer2(lines []string) int {
	totalScore := 0

	for _, line := range lines {
		var opponentMove, yourMove Move
		if i, err := ParseMove(rune(line[0])); err == nil {
			opponentMove = i
		}

		if i, err := ParseResult(rune(line[2])); err == nil {
			yourMove = CalculateMove(opponentMove, i)
		}

		totalScore += int(yourMove) + int(CalculateScore(opponentMove, yourMove))
	}

	return totalScore
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score ", Answer1(lines))
	fmt.Println("Total score for strategy: ", Answer2(lines))
}
