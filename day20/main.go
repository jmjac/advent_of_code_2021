package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x, y int
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	algo := scanner.Text()
	grid := make(map[Pos]int)

	scanner.Scan()
	var j int
	for scanner.Scan() {
		for i, k := range scanner.Text() {
			if k == '#' {
				grid[Pos{i, j}] = 1
			}
		}
		j++
	}

	fmt.Println("Part 1:", simulate(algo, grid, 2))
	fmt.Println("Part 2:", simulate(algo, grid, 50))
}

func simulate(algo string, grid map[Pos]int, steps int) int {
	var minX, minY, maxX, maxY int
	for i := 0; i < steps; i++ {
		newGrid := make(map[Pos]int)
		for pos := range grid {
			if pos.x > maxX {
				maxX = pos.x
			}
			if pos.x < minX {
				minX = pos.x
			}
			if pos.y > maxY {
				maxY = pos.y
			}
			if pos.y < minY {
				minY = pos.y
			}
		}

		cond := algo[0] == '#' && i%2 == 1
		for y := minY - 3; y <= maxY+2; y++ {
			for x := minX - 3; x <= maxX+2; x++ {
				if algo[convert(x, y, grid, cond, minX, minY, maxX, maxY)] == '#' {
					newGrid[Pos{x, y}] = 1
				}
			}
		}

		grid = newGrid
	}
	return len(grid)
}

func convert(x, y int, grid map[Pos]int, cond bool, minX, minY, maxX, maxY int) int {
	total := 0
	power := 256
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if i < minX || i > maxX || j < minY || j > maxY {
				if cond {
					total += power
				}
			} else {
				total += grid[Pos{i, j}] * power
			}
			power /= 2
		}
	}
	return total
}
