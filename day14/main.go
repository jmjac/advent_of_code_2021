package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	var polymer string
	rules := make(map[string]string, 0)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	polymer = scanner.Text()
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		line := strings.Split(scanner.Text(), " -> ")
		rules[line[0]] = line[1]
	}

	fmt.Println("Part 1:", simluate(polymer, rules, 10))
	fmt.Println("Part 2:", simluate(polymer, rules, 40))
}

func simluate(pol string, rules map[string]string, steps int) int {
	polymer := make(map[string]int)
	for i := 0; i < len(pol)-1; i++ {
		polymer[pol[i:i+2]] += 1
	}

	for s := 0; s < steps; s++ {
		newPol := make(map[string]int)
		for k, v := range polymer {
			newPol[string(k[0])+rules[k]] += v
			newPol[rules[k]+string(k[1])] += v
		}

		polymer = newPol
	}
	counter := make(map[rune]int)
	counter[rune(pol[0])] += 1
	counter[rune(pol[len(pol)-1])] += 1
	for k, v := range polymer {
		counter[rune(k[0])] += v
		counter[rune(k[1])] += v
	}
	most := 0
	least := polymer[pol[0:2]] * 10000
	for _, v := range counter {
		if v > most {
			most = v
		} else if v < least {
			least = v
		}

	}
	return most/2 - least/2
}
