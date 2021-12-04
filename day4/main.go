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
	var numbers []int
	var boards [][5][5]int
	scanner.Scan()
	for _, i := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(i)
		numbers = append(numbers, num)
	}

	numBoards := 0
	i := -1
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			numBoards++
			boards = append(boards, [5][5]int{})
			y = 0
			i++
			continue
		}
		for x, v := range strings.Fields(line) {
			num, _ := strconv.Atoi(v)
			boards[i][y][x] = num
		}
		y++
	}

	fmt.Println("Part 1:", findBoard(boards, numbers, true))
	fmt.Println("Part 2:", findBoard(boards, numbers, false))
}

func findBoard(boards [][5][5]int, numbers []int, findFirst bool) int {
	for num, drawn := range numbers {
		for i := 0; i < len(boards); i++ {
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if boards[i][y][x] == drawn {
						boards[i][y][x] = -1
					}
				}
			}
			if num >= 5 && checkWinning(boards[i]) {
				if findFirst || len(boards) == 1 {
					sum := 0
					for _, row := range boards[i] {
						for _, i := range row {
							if i != -1 {
								sum += i
							}
						}
					}
					return sum * drawn
				} else {
					boards[i] = boards[len(boards)-1]
					boards = boards[:len(boards)-1]
					i--
				}
			}
		}

	}
	return 0
}

func checkWinning(board [5][5]int) bool {
	for y := 0; y < 5; y++ {
		row := 0
		col := 0
		for x := 0; x < 5; x++ {
			row += board[y][x]
			col += board[x][y]
		}
		if row == -5 || col == -5 {
			return true
		}
	}
	return false
}
