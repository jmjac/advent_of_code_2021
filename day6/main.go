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

	fish := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _, i := range strings.Split(scanner.Text(), ",") {
			num, _ := strconv.Atoi(i)
			fish[num] += 1
		}
	}
	fmt.Println("Part 1:", simulate(fish, 80))
	fmt.Println("Part 2:", simulate(fish, 256))
}

func simulate(fish map[int]int, days int) int {
	for i := 0; i < days; i++ {
		newFish := make(map[int]int)
		for i := 8; i >= 1; i-- {
			newFish[i-1] = fish[i]
		}
		newFish[8] = fish[0]
		newFish[6] += fish[0]
		fish = newFish
	}
	count := 0
	for _, v := range fish {
		count += v
	}
	return count
}
