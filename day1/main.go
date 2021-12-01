package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	data := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		data = append(data, num)
	}

	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data []int) int {
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}
	return count

}

func part2(data []int) int {
	count := 0
	for i := 3; i < len(data); i++ {
		if data[i]+data[i-1]+data[i-2] > data[i-1]+data[i-2]+data[i-3] {
			count++
		}
	}
	return count

}
