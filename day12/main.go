package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Graph struct {
	edges map[string][]string
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	g := Graph{make(map[string][]string)}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		if _, ok := g.edges[line[0]]; !ok {
			g.edges[line[0]] = make([]string, 0)
		}
		if _, ok := g.edges[line[1]]; !ok {
			g.edges[line[1]] = make([]string, 0)
		}
		g.edges[line[0]] = append(g.edges[line[0]], line[1])
		g.edges[line[1]] = append(g.edges[line[1]], line[0])
	}
	fmt.Println("Part 1:", part1(g, visited))
	fmt.Println("Part 2:", part1(g, visitedTwice))
}

func part1(g Graph, visited_func func(string, []string) bool) int {
	count := 0
	paths := make([][]string, 0)
	paths = append(paths, make([]string, 1))
	paths[0][0] = "start"

	for len(paths) != 0 {
		path := paths[0]
		paths = paths[1:]
		if path[len(path)-1] == "end" {
			count++
		} else {
			for _, i := range g.edges[path[len(path)-1]] {
				if !visited_func(i, path) {
					newPath := make([]string, len(path), len(path)+1)
					copy(newPath, path)
					paths = append(paths, newPath)
					paths[len(paths)-1] = append(paths[len(paths)-1], i)
				}
			}
		}
	}
	return count
}

func visited(node string, path []string) bool {
	if unicode.IsUpper(rune(node[0])) {
		return false
	}
	for _, p := range path {
		if p == node {
			return true
		}
	}
	return false
}

func visitedTwice(node string, path []string) bool {
	if unicode.IsUpper(rune(node[0])) {
		return false
	}
	if node == "start" {
		return true
	}
	counter := make(map[string]int, 0)
	counter[node] = 1
	twice := false
	for _, p := range path {
		if unicode.IsUpper(rune(p[0])) {
			continue
		}
		counter[p]++
		v := counter[p]
		if v > 2 || v == 2 && twice {
			return true
		} else if v == 2 {
			twice = true
		}
	}

	return false

}
