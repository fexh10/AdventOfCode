package main

import (
	"bufio"
	. "fmt"
	"os"
)

type coord struct {
	x, y int
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func findTrailhead2(lines []string, i, j int, value int) (int) {
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[0]) || int(lines[i][j] - '0') != value{
        return 0
    }

	n := int(lines[i][j] - '0')

	if n == 9 {
		return 1
	}

	sum := 0

	for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
        sum += findTrailhead2(lines, i + dir[0], j + dir[1], n + 1)
    }
	return sum
}

func findTrailhead(lines []string, i, j int, value int, visited map[coord]bool) (int) {
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[0]) || visited[coord{i, j}] || int(lines[i][j] - '0') != value {
		return 0
	} else if int(lines[i][j] - '0') == 9 {
		return 1
	}
	sum := 0
	for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		ni, nj := i+dir[0], j+dir[1]
		temp := findTrailhead(lines, ni, nj, value + 1, visited)
		if temp > 0 {
			visited[coord{ni, nj}] = true
		}
		sum += temp
	}
	return sum
}

func part2(lines []string) (sum int) {
	for i, line := range lines {
		for j, r := range line {
			if r == '0' {
				sum += findTrailhead2(lines, i, j, 0)
			}
		}
	}
	return
}

func part1(lines []string) (sum int) {
	for i, line := range lines {
		for j, r := range line {
			if r == '0' {
				sum += findTrailhead(lines, i, j, 0, make(map[coord]bool))
			}
		}
	}
	return
}

func main() {
	var lines []string = input()
	Println(part1(lines))
	Println(part2(lines))
}
