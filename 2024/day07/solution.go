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

func parseInput() ([]uint64, [][]uint64) {
	lines := input()
	re := regexp.MustCompile(`\d+`)
	results := make([]uint64, 0)
	numbers := make([][]uint64, 0)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		n, _ := strconv.ParseUint(matches[0], 10, 64)
		results = append(results, n)
		temp := make([]uint64, 0)
		for i := 1; i < len(matches); i++ {
			n, _ = strconv.ParseUint(matches[i], 10, 64)
			temp = append(temp, n)
		}
		numbers = append(numbers, temp)
	}
	return results, numbers
}

func allCombinations(result, rem uint64, equation []uint64, part1 bool) bool {
	if len(equation) == 0 {
		return result == rem
	} 
	if part1 {
		return allCombinations(result, rem * equation[0], equation[1:], true) || allCombinations(result, rem + equation[0], equation[1:], true)
	} else {
		n, _ := strconv.ParseUint(strconv.FormatUint(rem, 10) + strconv.FormatUint(equation[0], 10), 10, 64)
		return allCombinations(result, rem * equation[0], equation[1:], false) || allCombinations(result, rem + equation[0], equation[1:], false) || allCombinations(result, n, equation[1:], false)

	}
}

func part2(results []uint64, numbers[][]uint64) (sum uint64) {
	for i, equation := range numbers {
		if allCombinations(results[i], 0, equation, false) {
			sum += 	results[i]
		}
	}
	return
}

func part1(results []uint64, numbers [][]uint64) (sum uint64) {
	for i, equation := range numbers {
		if allCombinations(results[i], 0, equation, true) {
			sum += 	results[i]
		}
	}
	return
}

func main() {
	results, numbers := parseInput()
	Println(part1(results, numbers))
	Println(part2(results, numbers))
}