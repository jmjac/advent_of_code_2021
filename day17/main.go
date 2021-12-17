package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	var sX, eX, sY, eY int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		x := line[2][2 : len(line[2])-1]
		y := line[3][2:]
		sX, _ = strconv.Atoi(strings.Split(x, "..")[0])
		eX, _ = strconv.Atoi(strings.Split(x, "..")[1])
		sY, _ = strconv.Atoi(strings.Split(y, "..")[0])
		eY, _ = strconv.Atoi(strings.Split(y, "..")[1])
	}

	p1, p2 := solve(sX, eX, sY, eY)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", p1, p2)
}

func solve(sX, eX, sY, eY int) (int, int) {
	best := 0
	count := 0
	for i := 0; i < 500; i++ {
		for j := -500; j < 500; j++ {
			vX := i
			vY := j
			x := 0
			y := 0
			bestTemp := 0
			for k := 0; k < 500; k++ {
				x += vX
				y += vY
				if vX > 0 {
					vX--
				} else if vX < 0 {
					vX++
				}
				vY--

				if y > bestTemp {
					bestTemp = y
				}

				if sX <= x && x <= eX && sY <= y && y <= eY {
					if bestTemp > best {
						best = bestTemp
					}
					count++
					break
				}
			}
		}
	}

	return best, count
}
