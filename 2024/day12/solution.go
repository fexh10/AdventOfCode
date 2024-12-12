package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

type pos struct {
	i, j int
}

type cell struct {
	let string
	val int
	dir int //0 = up, 1 = down, 2 = dx, 3 = sx
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func regions(lines []string, coord pos, r rune, checked map[pos]bool) (per int, area int) {
	if coord.i < 0 || coord.i >= len(lines) || coord.j < 0 || coord.j >= len(lines[0]) || rune(lines[coord.i][coord.j]) != r {
		return 1, 0
	} else if checked[coord] {
		return 0, 0
	}
	checked[coord] = true
	area += 1
	for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		ci, cj := coord.i+dir[0], coord.j+dir[1]
		p, a := regions(lines, pos{ci, cj}, r, checked)
		per += p
		area += a
	}
	return
}

func regions2(lines []string, coord pos, r rune, checked map[pos]bool, S map[cell][]int) (area int){
	checked[coord] = true
	area += 1
	if coord.j + 1 < len(lines[coord.i]) && !checked[pos{coord.i, coord.j + 1}] && r == rune(lines[coord.i][coord.j+1]) {
		regions2(lines, pos{coord.i, coord.j + 1}, r, checked, S)
	} else if !checked[pos{coord.i, coord.j + 1}] {
		S[cell{"j", coord.j + 1, 1}] = append(S[cell{"j", coord.j + 1, 1}], coord.i)
	}
	if coord.j - 1 >= 0 && !checked[pos{coord.i, coord.j - 1}] && r == rune(lines[coord.i][coord.j-1]) {
		regions2(lines, pos{coord.i, coord.j - 1}, r, checked, S)
	} else if !checked[pos{coord.i, coord.j - 1}] {
		S[cell{"j", coord.j, 0}] = append(S[cell{"j", coord.j, 0}], coord.i)
	}
	if coord.i + 1 < len(lines) && !checked[pos{coord.i + 1, coord.j}] && r == rune(lines[coord.i+1][coord.j]) {
		regions2(lines, pos{coord.i + 1, coord.j}, r, checked, S)
	} else if !checked[pos{coord.i + 1, coord.j}] {
		S[cell{"i", coord.i + 1, 3}] = append(S[cell{"i", coord.i + 1, 3}], coord.j)
	}
	if coord.i - 1 >= 0 && !checked[pos{coord.i - 1, coord.j}] && r == rune(lines[coord.i-1][coord.j]) {
		regions2(lines, pos{coord.i - 1, coord.j}, r, checked, S)
	} else if !checked[pos{coord.i - 1, coord.j}] {
		S[cell{"i", coord.i, 2}] = append(S[cell{"i", coord.i, 2}], coord.j)
	}
	return
}

func part2(lines []string) (sum int) {
	checked := make(map[pos]bool)
	for i, line := range lines {
		for j, r := range line {
			S := make(map[cell][]int)
			checked = make(map[pos]bool)
			a := regions2(lines, pos{i, j}, r, checked, S)
			L := 0
			for _, v := range S {
				if len(v) > 1 {
					sort.Ints(v)
					for k := 0; k < len(v)-1; k++ {
						if v[k]+1 != v[k+1] {
							L += 1
						}
					}
				}
				L += 1
			}
			sum += L * a
		}
	}
	return
}

func part1(lines []string) (sum int) {
	checked := make(map[pos]bool)

	for i, line := range lines {
		for j, r := range line {
			if !checked[pos{i, j}] {
				p, a := regions(lines, pos{i, j}, r, checked)
				sum += p * a
			}
		}
	}
	return
}

func main() {
	var lines []string = input()
	Println(part1(lines))
	Println(part2(lines))
}
