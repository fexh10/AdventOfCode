package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"strconv"
)

func input() string {
	sc := bufio.NewScanner(os.Stdin)
	var line string
	for sc.Scan() {
		line += sc.Text()
	}
	return line
}

func parseInput(lines string) map[string]int{
	stones := make(map[string]int)
	splitted := strings.Split(lines, " ")
	for _, e := range splitted {
		stones[e] += 1
	}
	return stones
}

func printStonesSum(stones map[string]int) (sum int) {
	for _, v := range stones {
		sum += v		
	}
	return
}

func main() {
	var line string = input()
	stones := parseInput(line)

	for i := 0; i < 75; i++ {
		temp := make(map[string]int)
		for key, value := range stones {
			if key == "0" {
				temp["1"] += value			
			} else if len(key) % 2 == 0 {
				temp[key[:len(key) / 2]] += value
				n, _ := strconv.Atoi(key[len(key) / 2:])
				temp[strconv.Itoa(n)] += value
			} else {
				n, _ := strconv.Atoi(key)
				temp[strconv.Itoa(n * 2024)] += value
			}
		}
		stones = temp
		if i == 24 {
			Println("Part 1:", printStonesSum(stones))
		}
	}
	Println("Part 2:", printStonesSum(stones))
}