package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func parte2(mappa map[Coordinate]int, folds [][]int) {
	for _, fold := range folds {
		for coord := range mappa{
			if fold[0] == 0 {
				if coord.x > fold[1] && mappa[coord] != 0 {
					strutt := Coordinate{fold[1] - (coord.x - fold[1] - 1) - 1, coord.y}
					mappa[strutt] = 1
					delete(mappa, coord)
				} 
			} else {
				if coord.y > fold[1] && mappa[coord] != 0 {
					strutt := Coordinate{coord.x, fold[1] * 2 - coord.y}
					mappa[strutt] = 1
					delete(mappa, coord)

				} 
			}
		}
	}
	printMap(mappa)
}

func printMap (mappa map[Coordinate]int) {
	maxX, maxY := 0, 0
	for coord := range mappa {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			coord := Coordinate{x, y}
			if mappa[coord] != 0 {
				Print("#")

			} else {
				Print(".")
			}
		}
		Println()
	}
}

func parte1(mappa map[Coordinate]int, folds [][]int) {
	for coord := range mappa {
		if folds[0][0] == 0 {
			if coord.x < folds[0][1] {
				strutt := Coordinate{folds[0][1] - (coord.x - folds[0][1] - 1) - 1, coord.y}
				mappa[strutt] += 1
				mappa[coord] = 0
			}
		} else {
			if coord.y > folds[0][1] {
				strutt := Coordinate{coord.x, folds[0][1] * 2 - coord.y}
				mappa[strutt] += 1
				mappa[coord] = 0
			}
		}
	}
	cont := 0
	for _, value := range mappa {
		if value != 0 {
			cont += 1
		}
	}
	Println(cont)
}

func parseInput() (mappa map[Coordinate]int, folds [][]int) {
	mappa = make(map[Coordinate]int)
	lines := input()
	for _, linea := range lines {
		if strings.Contains(linea, ",") {
			numeri := strings.Split(linea, ",")
			n1, _ := strconv.Atoi(numeri[0])
			n2, _ := strconv.Atoi(numeri[1])
			c := Coordinate{n1, n2}
			mappa[c] += 1
		} else if strings.Contains(linea, "fold") {
			stringa := strings.Split(linea, "=")
			str1 := string(stringa[0][len(stringa[0])-1])
			var n1 int
			if str1 == "x" {
				n1 = 0
			} else {
				n1 = 1
			}
			n2, _ := strconv.Atoi(stringa[1])
			var a []int
			a = append(a, n1, n2)
			folds = append(folds, a)
		}
	}
	return
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
	mappa, folds := parseInput()
	parte1(mappa, folds)
	parte2(mappa, folds)
}
