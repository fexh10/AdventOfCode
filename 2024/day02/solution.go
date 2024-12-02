package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func isSafe(numbers []string) (bool) {
	var sign int = 1
	for i := 0; i < len(numbers)-1; i++ {
		n1, _ := strconv.Atoi(numbers[i])
		n2, _ := strconv.Atoi(numbers[i+1])
		if i == 0 && n1 < n2 {
			sign = -1
		}
		diff := (n1 - n2) * sign
		if (diff < 1 || diff > 3){
			return false
		}
	}
	return true
}

func part2(input []string) int {
	var sum = 0
	for _, line := range input {
		numbers := strings.Split(line, " ")
		if isSafe(numbers) {
			sum += 1
		} else {
			for i := 0; i < len(numbers); i++ {
				temp := make([]string, len(numbers)-1)
				copy(temp, numbers[:i])
				copy(temp[i:], numbers[i+1:])
				if isSafe(temp) {
					sum += 1
					break
				}	
			}
		}
		
	}
	return sum
}

func part1(input []string) int {
	var sum = 0
	for _, line := range input {
		numbers := strings.Split(line, " ")
		if isSafe(numbers) {
			sum += 1
		}
	}
	return sum
}

func main() {
	var lines []string = input()
	Println(part1(lines))
	Println(part2(lines))
}
