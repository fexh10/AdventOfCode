package main

import (
	"bufio"
	. "fmt"
	"os"
)

type location struct {
	x int
	y int
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func parseInput(lines []string) map[string][]location {
	antennas := make(map[string][]location)
	for i, line := range lines {
		for j, c := range line {
			if c != '.' {
				antennas[string(c)] = append(antennas[string(c)], location{i, j})
			}
		}
	}
	return antennas
}

func part2(lines []string, antennas map[string][]location) int {
	antinodes := make(map[location]int)
	rows := len(lines)
	cols := len(lines[0])

	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				ax, ay := locations[i].x, locations[i].y
				bx, by := locations[j].x, locations[j].y
				
				cx, cy := ax, ay
				for cx >= 0 && cx < rows && cy >= 0 && cy < cols {
					antinodes[location{cx, cy}] = 1					
					cx -= bx - ax
					cy -= by - ay
				}

				dx, dy := bx, by
				for dx >= 0 && dx < rows && dy >= 0 && dy < cols {
					antinodes[location{dx, dy}] = 1
					dx += bx - ax
					dy += by - ay
				}
			}
		}
	}
	return len(antinodes)
}

func part1(lines []string, antennas map[string][]location) int {
	antinodes := make(map[location]int)
	rows := len(lines)
	cols := len(lines[0])

	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				dx := locations[j].x - locations[i].x
				dy := locations[j].y - locations[i].y

				newLocations := []location{
					{locations[i].x - dx, locations[i].y - dy},
					{locations[j].x + dx, locations[j].y + dy},
				}

				for _, newLoc := range newLocations {
					if newLoc.x >= 0 && newLoc.x < rows && newLoc.y >= 0 && newLoc.y < cols {
						antinodes[newLoc] += 1
					}
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	lines := input()
	antennas := parseInput(lines)
	Println(part1(lines, antennas))
	Println(part2(lines, antennas))
}
