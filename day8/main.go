package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	data := make([][]string, 0)
	outputs := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		data = append(data, make([]string, 0))
		outputs = append(outputs, make([]string, 0))
		line := strings.Split(scanner.Text(), " | ")

		for _, p := range strings.Split(line[0], " ") {
			p := SortString(p)
			data[i] = append(data[i], p)
		}

		for _, o := range strings.Split(line[1], " ") {
			o := SortString(o)
			outputs[i] = append(outputs[i], o)
		}
		i++

	}
	fmt.Println("Part 1:", part1(outputs))
	fmt.Println("Part 2:", part2(data, outputs))
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func part1(outputs [][]string) int {
	count := make(map[int]int)
	for _, i := range outputs {
		for _, k := range i {
			switch len(k) {
			case 2:
				count[1] += 1
			case 3:
				count[7] += 1
			case 4:
				count[4] += 1
			case 7:
				count[8] += 1
			}
		}
	}
	return count[1] + count[7] + count[4] + count[8]
}

func part2(data [][]string, outputs [][]string) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		encoding := make(map[int]string)
		length5 := make([]string, 3)
		length6 := make([]string, 3)

		for _, k := range data[i] {
			switch len(k) {
			case 2:
				encoding[1] = k
			case 3:
				encoding[7] = k
			case 4:
				encoding[4] = k
			case 7:
				encoding[8] = k
			case 5:
				length5 = append(length5, k)
			case 6:
				length6 = append(length6, k)

			}
		}

		for _, k := range length5 {
			if compare(k, encoding[1]) == 3 {
				encoding[3] = k
			} else if compare(k, encoding[4]) == 2 && compare(k, encoding[1]) == 4 {
				encoding[5] = k
			} else if compare(k, encoding[1]) == 4 && compare(k, encoding[4]) == 3 {
				encoding[2] = k
			}
		}
		for _, k := range length6 {
			if compare(encoding[4], k) == 0 {
				encoding[9] = k
			} else if compare(k, encoding[7]) == 3 {
				encoding[0] = k
			} else if compare(encoding[4], k) == 1 {
				encoding[6] = k
			}

		}

		decoding := make(map[string]int)
		for k, v := range encoding {
			decoding[v] = k
		}

		for i, z := range outputs[i] {
			sum += decoding[z] * int(math.Pow10(3-i))
		}
	}
	return sum
}

func compare(x, y string) int {
	if x == "" {
		return 0
	}
	count := 0
	for _, i := range x {
		found := false
		for _, j := range y {
			if i == j {
				found = true
				break
			}
		}
		if found == false {
			count++
		}
	}
	return count
}
