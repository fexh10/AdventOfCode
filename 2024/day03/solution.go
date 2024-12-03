package main

import (
	"bufio"
	. "fmt"
	"os"
	"regexp"
	"strconv"
)

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func sumMul(line string) (sum int) {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regex.FindAllStringSubmatch(line, -1)
	Println(matches)
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		sum += n1 * n2
	}
	return
}

func part2(input []string) int {
	var do bool = true
	var final string
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if i+3 < len(line) && line[i:i+4] == "do()" {
				do = true
				i += 3
			} else if i+6 < len(line) && line[i:i+7] == "don't()" {
				do = false
				i += 6
			} else if do {
				final += string(line[i])
			}
		}
	}
	return sumMul(final)
}

func part1(input []string) (sum int) {
	for _, line := range input {
		sum += sumMul(line)
	}
	return
}

func main() {
	var lines []string = input()
	Println(part1(lines))
	Println(part2(lines))
}