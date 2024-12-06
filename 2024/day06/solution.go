package main

import (
	"bufio"
	. "fmt"
	"os"
)

type position struct {
	x      int
	y      int
	direction string
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func findGuardPos(lines []string) (int, int) {
	for i, line := range lines {
		for j, c := range line {
			if c == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isLooping(lines []string) bool {
	rows := len(lines)
	cols := len(lines[0])
	x, y := findGuardPos(lines)
	guard := position{x, y, "up"}
	allPositions := make(map[position]int)
	for {
		allPositions[guard] += 1
		if v := allPositions[guard]; v > 2 {
			return true
		}
		switch guard.direction {
		case "up":
			if guard.x-1 >= 0 && (lines[guard.x-1][guard.y] == '.' || lines[guard.x-1][guard.y] == '^') {
				guard.x -= 1
			} else if guard.x-1 >= 0 && lines[guard.x-1][guard.y] == '#' {
				guard.direction = "right"
			} else {
				return false
			}
		case "down":
			if guard.x+1 < rows && (lines[guard.x+1][guard.y] == '.' || lines[guard.x+1][guard.y] == '^') {
				guard.x += 1
			} else if guard.x+1 < rows && lines[guard.x+1][guard.y] == '#' {
				guard.direction = "left"
			} else {
				return false
			}
		case "right":
			if guard.y+1 < cols && (lines[guard.x][guard.y+1] == '.' || lines[guard.x][guard.y+1] == '^') {
				guard.y += 1
			} else if guard.y+1 < cols && lines[guard.x][guard.y+1] == '#' {
				guard.direction = "down"
			} else {
				return false
			}
		case "left":
			if guard.y-1 >= 0 && (lines[guard.x][guard.y-1] == '.' || lines[guard.x][guard.y-1] == '^') {
				guard.y -= 1
			} else if guard.y-1 >= 0 && lines[guard.x][guard.y-1] == '#' {
				guard.direction = "up"
			} else {
				return false
			}
		}
	}
}

func part2(lines []string, allPos map[position]int) (sum int) {
	for p := range allPos {
		i, j := p.x, p.y
		if lines[i][j] == '.' {
			lineRunes := make([]rune, len(lines[i]))
			copy(lineRunes, []rune(lines[i]))
			lineRunes[j] = '#'
			newLines := make([]string, len(lines))
			copy(newLines, lines)
			newLines[i] = string(lineRunes)
			if isLooping(newLines) {
				sum += 1
			}
		}
	}
	return
}

func part1(lines []string) (map[position]int, int) {
	rows := len(lines)
	cols := len(lines[0])
	x, y := findGuardPos(lines)
	guard := position{x, y, "up"}
	allPositions := make(map[position]int)
	for {
		allPositions[position{guard.x, guard.y, ""}] += 1
		switch guard.direction {
		case "up":
			if guard.x-1 >= 0 && (lines[guard.x-1][guard.y] == '.' || lines[guard.x-1][guard.y] == '^') {
				guard.x -= 1
			} else if guard.x-1 >= 0 && guard.y+1 < cols && lines[guard.x-1][guard.y] == '#' {
				guard.direction = "right"
			} else {
				return allPositions, len(allPositions)
			}
		case "down":
			if guard.x+1 < rows && (lines[guard.x+1][guard.y] == '.' || lines[guard.x+1][guard.y] == '^') {
				guard.x += 1
			} else if guard.y-1 >= 0 && guard.x+1 < rows && lines[guard.x+1][guard.y] == '#' {
				guard.direction = "left"
			} else {
				return allPositions, len(allPositions)
			}
		case "right":
			if guard.y+1 < cols && (lines[guard.x][guard.y+1] == '.' || lines[guard.x][guard.y+1] == '^') {
				guard.y += 1
			} else if guard.x+1 < rows && guard.y+1 < cols && lines[guard.x][guard.y+1] == '#' {
				guard.direction = "down"
			} else {
				return allPositions, len(allPositions)
			}
		case "left":
			if guard.y-1 >= 0 && (lines[guard.x][guard.y-1] == '.' || lines[guard.x][guard.y-1] == '^') {
				guard.y -= 1
			} else if guard.x-1 >= 0 && guard.y-1 >= 0 && lines[guard.x][guard.y-1] == '#' {
				guard.direction = "up"
			} else {
				return allPositions, len(allPositions)
			}
		}
	}
}

func main() {
	var lines []string = input()
	allPos, part1 := part1(lines)
	Println(part1)
	Println(part2(lines, allPos))
}
