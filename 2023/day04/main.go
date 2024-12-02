package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func parseInput(lines []string, part1 bool) {
	var sum, tot int
	var copie []int 

	for i, line := range lines {	

		line = strings.ReplaceAll(line, "Card", "")
		line = strings.Replace(line, strconv.Itoa(i + 1), "", 1)
		line = strings.ReplaceAll(line, ":", "")
		line += " "
		numeri := strings.Split(line, "|")

		var vincenti, my_numbers []int
		var n string
		
		for _, char := range numeri[0] {
			if unicode.IsDigit(char) {
				n += string(char)
			} else {								
				temp, _ := strconv.Atoi(n)
				if temp != 0 {
					vincenti = append(vincenti, temp)
				}
				
				
				n = ""
			}
		}

		for _, char := range numeri[1] {
			if unicode.IsDigit(char) {
				n += string(char)
			} else {								
				temp, _ := strconv.Atoi(n)
				if temp != 0 {
					my_numbers = append(my_numbers, temp)
				}
				n = ""
			}
		}

		var cont int
		
		for _, n := range my_numbers {
			for _, vinc := range vincenti {
				if n == vinc {
					cont += 1
				}
			}
		}
		if part1 {
			sum += int(math.Pow(2, float64(cont) - 1))
		} else {
				if i > 1 {
					for j := 1; j <= cont; j++ {
						copie = append(copie, i + j)
					}
				}
				
			}
			
			
		
	}
	if part1 {
		Println(sum)
	} else {
		Println(tot)
	}
}

func parte2(lines []string) {
	parseInput(lines, false)
}

func parte1(lines []string) {
	parseInput(lines, true)
}

func main() {
	var lines[]string = input()
	parte1(lines)
	parte2(lines)
}