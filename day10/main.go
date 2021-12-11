package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("input.txt")

	data := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	p1, p2 := parts(data)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", p1, p2)
}

func parts(data []string) (part1, part2 int) {
	scores := make([]int, 0)
	endings := map[rune]rune{'}': '{', ']': '[', ')': '(', '>': '<'}
	scoringPart1 := map[rune]int{'}': 1197, ']': 57, ')': 3, '>': 25137}
	scoringPart2 := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

	for _, line := range data {
		s := make(stack, 0)
		var v rune
		for _, i := range line {
			if i == '}' || i == '>' || i == ')' || i == ']' {
				s, v = s.Pop()
				if endings[i] != v {
					part1 += scoringPart1[i]
					s = make(stack, 0)
					break
				}

			} else {
				s = s.Push(i)
			}
		}
		// Part2
		var score int
		for len(s) != 0 {
			s, v = s.Pop()
			score = score*5 + scoringPart2[v]
		}
		if score != 0 {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	part2 = scores[len(scores)/2]
	return
}

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}
