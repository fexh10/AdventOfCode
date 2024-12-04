package main

import (
	"bufio"
	. "fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
	"strings"
)

func input() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func part2(lines []string) (sum int) {
	spelled := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	for _, line := range lines {
		var first, last int
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if first == 0 {
					first = int(line[i] - '0')
				}
				last = int(line[i] - '0')
			}
			for k, v := range spelled {
				if strings.HasPrefix(line[i:], k) {
					if first == 0 {
						first = v
					}
					last = v
				}
			}
		}
		sum += first*10 + last
	}
	return
}

func part1(lines []string) (sum int) {
	re := regexp.MustCompile(`\d`)
	for _, line := range lines {
		matched := re.FindAllString(line, -1)
		addDigits := matched[0] + matched[len(matched)-1]
		n, _ := strconv.Atoi(addDigits)
		sum += n
	}
	return
}

func main() {
	var lines []string = input()
	Println(part1(lines))
	Println(part2(lines))
}
