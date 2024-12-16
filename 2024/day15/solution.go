package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func parseInput(lines []string) ([][]int, []pos) {
	var matrix [][]int
	var moves []pos
	for _, line := range lines {
		switch {
		case strings.Contains(line, "#"):
			var temp []int
			for _, r := range line {
				switch r {
				case '#':
					temp = append(temp, 0)
				case '.':
					temp = append(temp, -1)
				case 'O':
					temp = append(temp, 1)
				case '@':
					temp = append(temp, 2)
				}
			}
			matrix = append(matrix, temp)
		case strings.ContainsAny(line, "^v<>"):
			for _, r := range line {
				switch r {
				case '>':
					moves = append(moves, pos{0, 1})
				case '<':
					moves = append(moves, pos{0, -1})
				case '^':
					moves = append(moves, pos{-1, 0})
				case 'v':
					moves = append(moves, pos{1, 0})
				}
			}
		}
	}
	return matrix, moves
}

func findRobot(matrix [][]int) (int, int) {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 2 {
				return i, j
			}
		}
	}
	return -1, -1
}

func moveRobot(matrix [][]int, move pos, i, j int, r int) {
	ni, nj := i + move.i, j + move.j
	if matrix[ni][nj] == 0 {
		return
	}
	if matrix[ni][nj] == 1 {
		moveRobot(matrix, move, ni, nj, 1)
	} 
	if matrix[ni][nj] == -1 {
		matrix[ni][nj], matrix[i][j] = r, -1
	} 
}

func part1(matrix [][]int, moves []pos) (sum int) {
	for _, move := range moves {
		i, j := findRobot(matrix)
		moveRobot(matrix, move, i, j, 2)
	}
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 1 {
				sum += 100 * i + j
			}
		}
	}
	return
}

func main() {
	var lines []string = input()
	matrix, moves := parseInput(lines)
	Println(part1(matrix, moves))
}
