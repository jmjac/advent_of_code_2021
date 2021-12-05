package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Line struct {
	x1, y1, x2, y2 int
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	lines := make([]Line, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		x1, _ := strconv.Atoi(strings.Split(line[0], ",")[0])
		y1, _ := strconv.Atoi(strings.Split(line[0], ",")[1])
		x2, _ := strconv.Atoi(strings.Split(line[1], ",")[0])
		y2, _ := strconv.Atoi(strings.Split(line[1], ",")[1])
		lines = append(lines, Line{x1: x1, y1: y1, x2: x2, y2: y2})
	}

	now := time.Now()
	fmt.Println("Part 1:", countOverlaps(lines, false))
	fmt.Println("Part 2:", countOverlaps(lines, true))
	fmt.Println(time.Since(now))
}

func countOverlaps(lines []Line, diagonalLines bool) int {
	overlaps := make(map[cords]int)
	for _, line := range lines {
		if line.x1 == line.x2 {
			for i := min(line.y1, line.y2); i <= max(line.y1, line.y2); i++ {
				overlaps[cords{line.x1, i}] += 1
			}
		} else if line.y1 == line.y2 {
			for i := min(line.x1, line.x2); i <= max(line.x1, line.x2); i++ {
				overlaps[cords{i, line.y1}] += 1
			}
		} else if diagonalLines {
			for i := 0; i <= abs(line.x1-line.x2); i++ {
				overlaps[cords{line.x1 + i*(line.x2-line.x1)/abs(line.x2-line.x1), line.y1 + i*(line.y2-line.y1)/abs(line.y2-line.y1)}] += 1
			}
		}
	}

	var count int
	for _, v := range overlaps {
		if v >= 2 {
			count++
		}

	}
	return count
}

func max(a int, args ...int) int {
	max := a
	for _, i := range args {
		if i > max {
			max = i
		}
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type cords struct {
	x, y int
}
