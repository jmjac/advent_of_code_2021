package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	data := make([][]int, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, make([]int, 0))
		for _, k := range line {
			num, _ := strconv.Atoi(string(k))
			data[i] = append(data[i], num)
		}
		i++
	}
	flashes, sync := simulate(data, 100)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", flashes, sync)
}

type pos struct {
	y, x int
}

func simulate(data [][]int, daysLimit int) (dayCount, syncDay int) {
	var synced bool
	var days int
	var flashes int
	for !synced {
		days++
		flashing := make(map[pos]bool)
		flashed := make(map[pos]bool)
		for y := 0; y < len(data); y++ {
			for x := 0; x < len(data[y]); x++ {
				data[y][x] += 1
				if data[y][x] == 10 {
					flashing[pos{y, x}] = true
				}

			}
		}

		for len(flashing) != 0 {
			newFlashing := make(map[pos]bool)
			for k := range flashing {
				if flashed[k] {
					continue
				}
				flashed[k] = true
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if k.y+y < 0 || k.y+y >= len(data) || k.x+x < 0 || k.x+x >= len(data[0]) || x == 0 && y == 0 {
							continue
						}
						data[k.y+y][k.x+x] += 1
						if data[k.y+y][k.x+x] >= 10 {
							newFlashing[pos{k.y + y, k.x + x}] = true
						}
					}
				}

			}
			flashing = newFlashing
		}

		for k := range flashed {
			data[k.y][k.x] = 0
			flashes++
		}

		if days == daysLimit {
			dayCount = flashes
		}
		synced = len(flashed) == len(data)*len(data[0])
	}
	syncDay = days
	return

}
