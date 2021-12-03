package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	data := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())

	}
	mostBin, leastBin := findMostandLeast(data)
	mostInt, _ := strconv.ParseInt(mostBin, 2, 0)
	leastInt, _ := strconv.ParseInt(leastBin, 2, 0)
	fmt.Println("Part 1:", mostInt*leastInt)
	mostFiltered, _ := strconv.ParseInt(filter(data, 0, true), 2, 0)
	leastFiltered, _ := strconv.ParseInt(filter(data, 0, false), 2, 0)
	fmt.Println("Part 2:", mostFiltered*leastFiltered)
}

func findMostandLeast(data []string) (string, string) {
	bits := make([]int, len(data[0]))
	for _, i := range data {
		for k, j := range i {
			if j == '1' {
				bits[k] += 1
			}
		}
	}
	most := ""
	least := ""
	for _, i := range bits {
		if len(data)-i*2 > 0 {
			most += "1"
			least += "0"
		} else {
			most += "0"
			least += "1"
		}
	}
	return most, least
}

func filter(data []string, pos int, mostCommon bool) string {
	if len(data) == 1 || len(data[0]) <= pos {
		return data[0]
	}
	var one []string
	var zero []string
	count := 0
	for _, i := range data {
		if i[pos] == '1' {
			count += 1
			one = append(one, i)
		} else {
			zero = append(zero, i)
		}
	}
	if (mostCommon && count*2 >= len(data)) || (!mostCommon && count*2 < len(data)) {
		return filter(one, pos+1, mostCommon)
	}
	return filter(zero, pos+1, mostCommon)
}
