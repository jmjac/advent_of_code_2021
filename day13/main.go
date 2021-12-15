package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type split struct {
	dir string
	val int
}

func main() {
	f, _ := os.Open("input.txt")

	data := make(map[Pos]bool, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		data[Pos{x, y}] = true
	}

	splits := make([]split, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")[2]
		dir := strings.Split(line, "=")[0]
		val, _ := strconv.Atoi(strings.Split(line, "=")[1])
		splits = append(splits, split{dir, val})
	}

	fmt.Println("Part 1:", simluate(data, splits))
}

func simluate(paper map[Pos]bool, splits []split) (part1 int) {
	for i, s := range splits {
		if i == 1 {
			part1 = len(paper)
		}
		newPaper := make(map[Pos]bool)
		if s.dir == "x" {
			for pos := range paper {
				if pos.x > s.val {
					newPaper[Pos{2*s.val - pos.x, pos.y}] = true
				} else {
					newPaper[pos] = true
				}
			}
		} else {
			for pos := range paper {
				if pos.y > s.val {
					newPaper[Pos{pos.x, 2*s.val - pos.y}] = true
				} else {
					newPaper[pos] = true
				}
			}
		}
		paper = newPaper
	}
	fmt.Println("Part 2:")
	part2(paper)
	return
}

func part2(paper map[Pos]bool) {
	maxX, maxY := 0, 0
	for pos := range paper {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if paper[Pos{x, y}] {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}

}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
