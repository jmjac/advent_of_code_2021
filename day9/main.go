package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	data := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		data = append(data, make([]int, 0))
		for _, k := range scanner.Text() {
			num, _ := strconv.Atoi(string(k))
			data[i] = append(data[i], num)
		}
		i++
	}
	risk, lowPoints := part1(data)
	fmt.Println("Part 1:", risk)
	fmt.Println("Part 2:", part2(data, lowPoints))
}

type cords struct {
	y, x int
}

func part1(data [][]int) (int, []cords) {
	risk := 0
	lowPoints := make([]cords, 0)
	for i, row := range data {
		for k := 0; k < len(row); k++ {
			if (k == 0 && row[k] < row[k+1]) || (k == len(row)-1 && row[k] < row[k-1]) || (k != 0 && k != len(data)-1 && row[k] < row[k-1] && row[k] < row[k+1]) {
				if (i == 0 && row[k] < data[i+1][k]) || (i == len(data)-1 && row[k] < data[i-1][k]) || (i != 0 && i != len(data)-1 && row[k] < data[i-1][k] && row[k] < data[i+1][k]) {
					risk += row[k] + 1
					lowPoints = append(lowPoints, cords{i, k})
				}
			}
		}
	}
	return risk, lowPoints
}

func part2(data [][]int, lowPoints []cords) int {
	basins := make([]int, 0, len(lowPoints))
	for _, cord := range lowPoints {
		checked := make(map[cords]bool)
		y, x := cord.y, cord.x
		basins = append(basins, expandBasin(y, x, data, checked))

	}
	sort.Ints(basins)
	sum := 1
	for i := len(basins) - 1; i > len(basins)-4; i-- {
		sum *= basins[i]
	}
	return sum
}

func expandBasin(y, x int, data [][]int, checked map[cords]bool) int {
	if x < 0 || y < 0 || y >= len(data) || x >= len(data[0]) {
		return 0
	}
	if _, ok := checked[cords{y, x}]; ok || data[y][x] == 9 {
		return 0
	}
	checked[cords{y, x}] = true
	return expandBasin(y-1, x, data, checked) + expandBasin(y+1, x, data, checked) + expandBasin(y, x-1, data, checked) + expandBasin(y, x+1, data, checked) + 1
}
