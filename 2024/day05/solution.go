package main

import (
	"bufio"
	. "fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(lines []string) (map[int][]int, [][]int) {
	before := make(map[int][]int)
	update := make([][]int, 0)

	for _, line := range lines {
		splitted := strings.Split(line, "|")
		if line == "" {
			continue
		} else if len(splitted) == 1 {
			splitted = strings.Split(line, ",")
			ints := make([]int, len(splitted))
			for i, s := range splitted {
				ints[i], _ = strconv.Atoi(s)
			}
			update = append(update, ints)
		} else {
			first, _ := strconv.Atoi(splitted[0])
			second, _ := strconv.Atoi(splitted[1])
			before[second] = append(before[second], first)
		}
	}
	return before, update
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func indexes(slice, update []int, k int) int {
	for _, item := range slice {
	  index := slices.Index(update, item)
	  if index != -1 && k < index {
		return index
	  }
	}
	return -1
  }

func is_correct(before map[int][]int, update []int) bool {
	for i, page := range update {
		if indexes(before[page], update, i) != -1 {
			return false
		}
	}
	return true
}

func part2(before map[int][]int, updates [][]int) (sum int) {
	for _, update := range updates {
		newslice := update
		if is_correct(before, newslice) {
			continue
		}
		for !is_correct(before, newslice) {
			for i, page := range newslice {
				if index := indexes(before[page], newslice, i); index != -1 {
					newslice[i], newslice[index] = newslice[index], newslice[i]
				}
			}
		}
		sum += newslice[len(newslice)/2]
	}
	return
}

func part1(before map[int][]int, updates [][]int) (sum int) {
	for _, update := range updates {
	  if is_correct(before, update) {
		sum += update[len(update)/2]
	  }
	}
	return
  }

func main() {
	before, update := parseInput(input())
	Println(part1(before, update))
	Println(part2(before, update))
}
