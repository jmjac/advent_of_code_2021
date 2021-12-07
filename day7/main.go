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

	data := make([]int, 0)
	scanner := bufio.NewScanner(f)
	max := 0
	for scanner.Scan() {
		for _, i := range strings.Split(scanner.Text(), ",") {
			num, _ := strconv.Atoi(i)
			data = append(data, num)
			if num > max {
				max = num
			}
		}
	}

	fmt.Println("Part 1:", part1(data, max, abs))
	sum := func(x int) int { return sumAll(abs(x)) }
	fmt.Println("Part 1:", part1(data, max, sum))
}

func part1(data []int, max int, sumMethod func(int) int) int {
	best := max * 100000
	for i := 0; i < max; i++ {
		sum := 0
		for _, v := range data {
			sum += sumMethod(i - v)
		}
		if best >= sum {
			best = sum
		}
	}
	return best
}

func sumAll(x int) int {
	return (x*x + x) / 2
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
