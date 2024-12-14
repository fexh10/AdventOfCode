package main

import (
	"bufio"
	. "fmt"
	"os"
	"regexp"
	"strconv"
)

type pos struct {
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

const rows, cols = 103, 101

func parseInput(lines []string) [][]pos {
	re := regexp.MustCompile(`-?\d+`)
	var res [][]pos
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		numbers := make([]int, len(matches))
		for i := 0; i < len(matches); i++ {
			numbers[i], _ = strconv.Atoi(matches[i])
		}
		res = append(res, []pos{{numbers[0], numbers[1]}, {numbers[2], numbers[3]}})
	}
	return res
}

func getWrappedPos(x, y, vx, vy, time int) pos {
	finalX := (x + vx*time) % cols
	if finalX < 0 {
		finalX += cols
	}

	finalY := (y + vy*time) % rows
	if finalY < 0 {
		finalY += rows
	}
	return pos{finalX, finalY}
}

func findTree(positions map[pos]bool) bool {
	for k := range positions {
		// check if there is a square of 3x3 with the current position as the center
		if positions[pos{k.x - 1, k.y - 1}] && positions[pos{k.x, k.y - 1}] && positions[pos{k.x + 1, k.y - 1}] &&
			positions[pos{k.x - 1, k.y}] && positions[pos{k.x, k.y}] && positions[pos{k.x + 1, k.y}] &&
			positions[pos{k.x - 1, k.y + 1}] && positions[pos{k.x, k.y + 1}] && positions[pos{k.x + 1, k.y + 1}] {
			return true
		}
	}
	return false
}

func part2(input [][]pos) (res int) {
	res = -1
	t := 1
	for {
		positions := make(map[pos]bool)
		for _, values := range input {
			x, y, vx, vy := values[0].x, values[0].y, values[1].x, values[1].y
			finalPos := getWrappedPos(x, y, vx, vy, t)
			positions[finalPos] = true
		}
		if findTree(positions) {
			res = t
			break
		}
		t += 1
	}
	return res
}

func part1(input [][]pos) int {
	time := 100
	quadCnt := []int{0, 0, 0, 0}

	for _, values := range input {
		x, y, vx, vy := values[0].x, values[0].y, values[1].x, values[1].y
		finalPos := getWrappedPos(x, y, vx, vy, time)
		if finalPos.x == (cols -1) / 2 || finalPos.y == (rows - 1) / 2 {
			continue
		}
		quadrant := 0
		if finalPos.x > (cols - 1) / 2 {
			quadrant += 1
		}
		if finalPos.y > (rows - 1) / 2 {
			quadrant += 2
		}
		quadCnt[quadrant] += 1
	}
	res := 1
	for _, cnt := range quadCnt {
		res *= cnt
	}
	return res
}

func main() {
	var lines []string = input()
	var input [][]pos = parseInput(lines)
	Println(part1(input))
	Println(part2(input))
}
