package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type points struct {
	x, y float64
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func parseLine(line string, re *regexp.Regexp) points {
	matches := re.FindAllString(line, -1)
	x, _ := strconv.ParseFloat(matches[0], 64)
	y, _ := strconv.ParseFloat(matches[1], 64)
	return points{x, y}
}

func parseInput(lines []string) ([]points, []points, []points) {
	var a, b, x []points
	re := regexp.MustCompile(`\d+`)
	for i := 0; i < len(lines); i += 4 {
		a = append(a, parseLine(lines[i], re))
		b = append(b, parseLine(lines[i+1], re))
		x = append(x, parseLine(lines[i+2], re))
	}
	return a, b, x
}

func part2(a, b, x []points) (sum int) {
	for i, point := range x {
		point.x += 10000000000000
		point.y += 10000000000000
		tolerance := 0.0001
		det := a[i].x * b[i].y - a[i].y * b[i].x
		c := ((point.x * b[i].y) + (point.y * -b[i].x)) / det
		d := (point.y - a[i].y * c) / b[i].y
		rc := math.Round(c)
		rd := math.Round(d)
		if math.Abs(c - rc) <= tolerance && math.Abs(d - rd) <= tolerance {
			sum += int(c*3 + d)
		}
	}
	return
}

func part1(a, b, x []points) (sum int) {
	for i, point := range x {
		det := a[i].x * b[i].y - a[i].y * b[i].x
		c := ((point.x * b[i].y) + (point.y * -b[i].x)) / det
		d := point.y - a[i].y * c
		if d / b[i].y == float64(int(d / b[i].y)) && c >= 0 && d / b[i].y >= 0 {
			sum += int(c * 3 + d / b[i].y)
		}
	}
	return
}

func main() {
	var lines []string = input()
	a, b, x := parseInput(lines)
	Println(part1(a, b, x))
	Println(part2(a, b, x))
}