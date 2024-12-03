package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
	"math"
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

func parseInput(line string, i int) (numbers, victoryNumbers []int){
	temp := "Card " + strconv.Itoa(i + 1) + ": "
	line = strings.ReplaceAll(line, temp, "")
	allNumbers := strings.Split(line, " | ")
	slice := strings.Split(allNumbers[0], " ")

	for _, number := range slice {
		n, err := strconv.Atoi(number)
		if err == nil {
			numbers = append(numbers, n)
		}
	}

	slice = strings.Split(allNumbers[1], " ")
	for _, number := range slice {
		n, err := strconv.Atoi(number)
		if err == nil {
			victoryNumbers = append(victoryNumbers, n)
		}
	}
	return
}

func winningNumbers(numbers, victoryNumbers []int) (winningNumbers int) {
	for _, number := range numbers {
		for _, victoryNumber := range victoryNumbers {
			if number == victoryNumber {
				winningNumbers += 1
			}
		}
	}
	return
}

func processCard(lines []string, i int) int {
	if i >= len(lines) {
		return 0
	}
	numbers, victoryNumbers := parseInput(lines[i], i)
	winningNumbers := winningNumbers(numbers, victoryNumbers)
	total := 1
	for j := 0; j < winningNumbers; j++ {
		total += processCard(lines, i + 1 + j)
	}
	return total
}

func part2(lines []string) (totalScratchcards int) {
	for i := range lines {
		totalScratchcards += processCard(lines, i)
	}
	return
}

func part1(lines []string) (sum int) {
	for i, line := range lines {
		numbers, victoryNumbers := parseInput(line, i)
		winningNumbers := winningNumbers(numbers, victoryNumbers)
		sum += int(math.Pow(2, float64(winningNumbers) - 1))
	}
	return
}

func main() {
	var lines[]string = input()
	Println(part1(lines))
	Println(part2(lines))
}