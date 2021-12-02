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
	data := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Split(line, " "))
	}

	fmt.Println("Part 1: ", part1(data))
	fmt.Println("Part 2: ", part2(data))
}

func part1(data [][]string) int {
	x := 0
	depth := 0
	for _, i := range data {
		v, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "forward":
			x += v
		case "up":
			depth -= v
		case "down":
			depth += v
		}
	}
	return x * depth
}

func part2(data [][]string) int {
	x := 0
	depth := 0
	aim := 0
	for _, i := range data {
		v, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "forward":
			x += v
			depth += v * aim
		case "up":
			aim -= v
		case "down":
			aim += v
		}
	}
	return x * depth
}
