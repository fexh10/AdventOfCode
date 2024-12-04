package main

import (
	. "fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
)

func parte1(lines []string) {
	mappa := make(map[int]int)
	//parsing input
	numeri_stringhe := strings.Split(lines[0], ",")
	for _, element := range numeri_stringhe {
		number, _ := strconv.Atoi(element)
		mappa[number] += 1
	}

	for i := 0; i < 256; i++ {
		newmap := make(map[int]int)
		for key, value := range mappa {
			if key != 0 {
				newmap[key - 1] += value
			} else {
				newmap[6] += value
				newmap[8] += value
			}
		}
		mappa = newmap
	}
	somma := 0
	for _, value := range mappa {
		somma +=  value
	}
	Println(somma)
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
	start := time.Now()
	lines := input()
	parte1(lines)
	Println(time.Since(start))
}