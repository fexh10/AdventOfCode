package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"unicode"
)

func adiacente(lines[] string, strIndexini, strInidexFin, nLine int) bool {
	for i, line := range lines {
		if i == nLine || i == nLine - 1 || i == nLine + 1 {
			for j, char := range line {
				if (j >= strIndexini  && j <= strInidexFin) || (j >= strIndexini + 1  && j <= strInidexFin + 1) || (j >= strIndexini - 1  && j <= strInidexFin - 1) {
					if !unicode.IsDigit(char) && char != '.' {
						return true
					}
				}
			}
		}
	}
	return false
}

func adiacente2(lines[] string, lineIndex, charIndex int) (bool) {
	for i, line := range lines {
		line += "."
		if i == lineIndex || i == lineIndex - 1 || i == lineIndex + 1 {
			for j, char := range line {
				if (j == charIndex) || (j == charIndex + 1) || (j == charIndex - 1) {
					if  unicode.IsDigit(char) {
						return true
					}
				}
			}
		}
	}
	return false
}

func trovaNumeri(lines[] string, indexLineStar, indexCharStar int) []int {
	var numeri []int
	var n string
	var indexIni, temp int

	for i, line := range lines {
		line += "."
		var linea, trovato bool = false, false
		for j, char := range line {
			if i == indexLineStar || i == indexLineStar - 1 || i == indexLineStar + 1 {
				linea = true
			} 
			if linea {
				if j == indexCharStar || j == indexCharStar - 1 || j == indexCharStar + 1 {
					if unicode.IsDigit(char) {
						trovato = true
					}
				}
				linea = false
			}

			if unicode.IsDigit(char) {
				if indexIni == 0 {
					indexIni = j
				}
				n += string(char)
			} else {
				if n != "" {
					
				temp, _ = strconv.Atoi(n)
				if trovato {
					numeri = append(numeri, temp)
					trovato = false
				}
						
					
					n = ""
					indexIni = 0
				}
			}

			
		}
	}
	return numeri
}

func parte2(lines[] string) {
	var sum int

	for i, line := range lines {
		line += "."
		for j, char := range line {
			if char == '*' {

				var ok bool

				ok = adiacente2(lines, i, j)
				if ok {
					numeri := trovaNumeri(lines, i, j)
					if len(numeri) != 1 {
						for i := 0; i < len(numeri); i++ {
							if i != len(numeri) - 1{
								sum += numeri[i] * numeri[i + 1]
							}
							
						}
					}
				}
			}
		}
	}
	Println(sum)
}

func parte1(lines[] string) {
	var sum int

	for i, line := range lines {
		var n string
		var indexIni int
		line += "."
		for j, char := range line {
			if unicode.IsDigit(char) {
				if indexIni == 0 {
					indexIni = j
				}
				n += string(char)
			} else {
				if n != "" {
					if adiacente(lines, indexIni, j - 1, i) {
						temp, _ := strconv.Atoi(n)
						sum += temp
					}
					n = ""
					indexIni = 0
				}
			}
		}
	}
	Println(sum)
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
	var lines[]string = input()
	parte1(lines)
	parte2(lines)
}