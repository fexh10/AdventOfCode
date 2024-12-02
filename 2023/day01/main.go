package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parte2(lines []string) {
	diz := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	var primo, ultimo string
	index := 0
	var numeri []int

	for _, line := range lines {

		for number := range diz {
			if strings.Contains(line, number) {
				if strings.Index(line, number) <= index {
					primo = diz[number]
				} else {
					ultimo = diz[number]
				}
				index = strings.Index(line, number)
			}
			for i, char := range line {
				for j := 0; j < 9; j++ {
					if char == rune('0'+j) {
						if i <= index {
							primo = string(char)
						} else {
							ultimo = string(char)
						}
						index = i
					}
				}
			}
		}

		if ultimo == "" {
			ultimo = primo
		}
		n, _ := strconv.Atoi(primo + ultimo)
		numeri = append(numeri, n)
		primo, ultimo = "", ""
		index = 0
	}
	sum := 0

	for _, n := range numeri {
		sum += n
	}
	Println(numeri[0])
}

func findNumbers(lines []string) []string {
	var first, last string
	var numeri []string

	for _, lin := range lines {
		for _, char := range lin {
			if unicode.IsDigit(char) {
				last = string(char)
				if first == "" {
					first = string(char)
				}
			}
		}
	
		temp := first + last
		numeri = append(numeri, temp)
		first, last = "", ""
	}
	return numeri
}

func parte1(lines []string) {
	var sum int

	numeri := findNumbers(lines)

	for _, e := range numeri {
		n, _ := strconv.Atoi(e)
		sum += n
	}
	Println(sum)
}

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

func main() {
	var lines []string = input()
	parte1(lines)
	parte2(lines)
}
