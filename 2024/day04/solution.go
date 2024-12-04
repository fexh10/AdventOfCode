package main

import (
	"bufio"
	. "fmt"
	"os"
)

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func findWors(lines []string, i, j int) (sum int) {
	// horizontal
	if j+4 < len(lines[i]) && lines[i][j:j+4] == "XMAS" {
		sum += 1
	}
	// horizontal backwards
	if j+3 < len(lines[i]) && lines[i][j:j+4] == "SAMX" {
		sum += 1
	}
	// vertical
	if i+3 < len(lines) && lines[i][j] == 'X' && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
		sum += 1
	}
	// vertical backwards
	if i+3 < len(lines) && lines[i][j] == 'S' && lines[i+1][j] == 'A' && lines[i+2][j] == 'M' && lines[i+3][j] == 'X' {
		sum += 1
	}
	// diagonal up left
	if i-3 >= 0 && j-3 >= 0 && lines[i][j] == 'X' && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
		sum += 1
	}
	// diagonal up right
	if i-3 >= 0 && j+3 < len(lines[i]) && lines[i][j] == 'X' && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
		sum += 1
	}
	// diagonal down left
	if i+3 < len(lines) && j-3 >= 0 && lines[i][j] == 'X' && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
		sum += 1
	}
	// diagonal down right
	if i+3 < len(lines) && j+3 < len(lines[i]) && lines[i][j] == 'X' && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
		sum += 1
	}
	return
}

func part2(lines []string) (sum int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if (i + 2 < len(lines) && j < len(lines[i]) - 2) {
				word1 := string(lines[i][j]) + string(lines[i + 1][j + 1]) + string(lines[i + 2][j + 2])
				word2 := string(lines[i][j + 2]) + string(lines[i + 1][j + 1]) + string(lines[i + 2][j])
				if  (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
					sum +=1
				}
			}
		}
	}
	return
}

func part1(lines []string) (sum int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			sum += findWors(lines, i, j)
		}
	}
	return
}

func main() {
	var lines []string = input()
	Println(part1(lines))
	Println(part2(lines))
}
