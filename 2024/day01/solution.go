package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func findMin(numbers *map[int]int) (n int) {
	n = math.MaxInt
	for key := range *numbers {
		if key < n {
			n = key
		}
	}
	(*numbers)[n] -= 1
	if (*numbers)[n] == 0 {
		delete(*numbers, n)
	}
	return
}

func parseInput(input []string) (leftList, rightList map[int]int) {
	leftList = make(map[int]int)
	rightList = make(map[int]int)
	for _, line := range input {
		splitted := strings.Split(line, "   ")
		n, _ := strconv.Atoi(splitted[0])
		leftList[n] += 1
		n, _ = strconv.Atoi(splitted[1])
		rightList[n] += 1
	}
	return
}

func part2(input []string) int {
	leftList, rightList := parseInput(input)
	var sum = 0
	for key := range leftList {
		if v, ok := rightList[key]; ok {
			sum += key * v
		}
		delete(leftList, key)
	}
	return sum
}

func part1(input []string) int {
	leftList, rightList := parseInput(input)	
	var sum = 0
	for len(leftList) > 0 {
		sum += int(math.Abs(float64(findMin(&leftList) - findMin(&rightList))))
	}
	return sum
}

func input() []string {
	file, _ := os.Open(os.Args[1] + ".txt")
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
	Println(part1(lines))
	Println(part2(lines))
}
