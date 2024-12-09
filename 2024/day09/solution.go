package main

import (
	"bufio"
	. "fmt"
	"os"
	"slices"
	"strconv"
)

func input() string {
	sc := bufio.NewScanner(os.Stdin)
	var lines string
	for sc.Scan() {
		lines = sc.Text()
	}
	return lines
}

func parseInput(line string) []int {
	var slice []int
	fileID := 0

	for i, c := range line {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				slice = append(slice, fileID)
			}
			fileID += 1
		} else {
			for j := 0; j < n; j++ {
				slice = append(slice, -1)
			}
		}
	}
	return slice
}

func part2(line string) (checksum int) {
	diskMap := make([][]int, 0)
	fileID := 0
	for i, c := range line {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 && n != 0 {
			temp := make([]int, 0)
			for j := 0; j < n; j++ {
				temp = append(temp, fileID)
			}
			diskMap = append(diskMap, temp)
			fileID += 1
		} else if n != 0 {
			temp := make([]int, 0)
			for j := 0; j < n; j++ {
				temp = append(temp, -1)
			}
			diskMap = append(diskMap, temp)
		}
	}
	for i := len(diskMap) - 1; i > 0; i-- {
		if diskMap[i][0] == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if diskMap[j][0] == -1 {
				if len(diskMap[i]) < len(diskMap[j]) {
					temp := diskMap[j][len(diskMap[i]):]
					copyslice := make([]int, len(diskMap[i]))
					copy(copyslice, diskMap[j])
					copy(diskMap[j], diskMap[i])
					diskMap[j] = diskMap[j][:len(diskMap[i])]
					copy(diskMap[i], copyslice)
					for w := j + 1; w < len(diskMap); w++ {
						temp, diskMap[w] = diskMap[w], temp
					}
					diskMap = append(diskMap, temp)
					i++
					break
				} else if len(diskMap[i]) == len(diskMap[j]) {
					diskMap[j], diskMap[i] = diskMap[i], diskMap[j]
					break
				}
			}
		}
	}

	index := 0
	for _, num := range diskMap {
		for _, n := range num {
			if n != -1 {
				checksum += index * n
			}
			index += 1
		}
	}
	return
}

func part1(diskMap []int) (checksum int) {
	for i := len(diskMap) - 1; i >= 0; i-- {
		index := slices.Index(diskMap, -1)
		if index == -1 {
			break
		}
		diskMap[index] = diskMap[i]
		diskMap = diskMap[:i]
	}

	for i, n := range diskMap {
		checksum += i * n
	}
	return
}

func main() {
	line := input()
	diskMap := parseInput(line)
	Println(part1(diskMap))
	Println(part2(line))
}
