package main

import (
	. "fmt"
	"os"
	"bufio"
	"strconv"
)

type coord struct {
	x, y int
}

func parte2(input []string, i, j int, mappa map[coord]bool) (size int) {
	struttura := coord{i, j}
	if mappa[struttura] || input[i][j] == byte('9') {
		return 0
	}
	mappa[struttura] = true
	size += 1
	if i != 0 {
		size += parte2(input, i - 1, j, mappa)
	}
	if j != 0 {
		size += parte2(input, i, j - 1, mappa)
	}
	if j != len(string(input[i])) - 1  {
		size += parte2(input, i, j + 1, mappa)
	}
	if i != len(input) - 1 {
		size += parte2(input, i  + 1, j, mappa)
	}
	return
}

func parte1(input []string) (cont int) {
	var slice[]int;
	for i, line := range input {
		for j, runa := range line {
			if i != 0 && runa >= rune(input[i - 1][j]) {
				continue
			}
			if j != 0 && runa >= rune(input[i][j - 1]) {
				continue
			}
			if j != len(line) - 1 && runa >= rune(input[i][j + 1]) {
				continue
			}
			if i != len(input) - 1 && runa >= rune(input[i + 1][j]) {
				continue
			}
			var mappa = make(map[coord]bool)
			slice = append(slice, parte2(input, i, j, mappa))
			temp, _ := strconv.Atoi(string(runa))
			cont += temp + 1;
		}
	}
	var max1, max2, max3 int;
	for _, n := range slice {
		if n > max1 {
			max2, max3 = max1, max2
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n
		}
	}
	Println(max1 * max2 * max3)
	return
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
	input := input()
	Println(parte1(input))
}