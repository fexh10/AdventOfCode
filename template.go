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

func main() {
	var lines []string = input()
}