package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	data := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		data = append(data, make([]int, 0))
		for _, k := range scanner.Text() {
			num, _ := strconv.Atoi(string(k))
			data[i] = append(data[i], num)
		}
		i++
	}

	bigBoard := big(data, 5)
	fmt.Println("Part 1:", leastRisk(data))
	fmt.Println("Part 2:", leastRisk(bigBoard))
}

type Pos struct {
	x, y int
}

func leastRisk(data [][]int) int {
	cost := make(map[Pos]int)
	q := make(priorityQueue, 0, 1000)
	q = append(q, item{Pos{0, 0}, 0})
	cost[Pos{0, 0}] = 0

	for len(q) != 0 {
		pos := heap.Pop(&q).(item).pos
		for _, i := range neighbours(pos) {
			if i.x < 0 || i.y < 0 || i.x >= len(data[0]) || i.y >= len(data) {
				continue
			}
			risk := cost[pos] + data[i.y][i.x]
			if c, ok := cost[i]; !ok || risk < c {
				cost[i] = risk
				heap.Push(&q, item{i, risk})
			}
		}
	}
	end := Pos{len(data[0]) - 1, len(data) - 1}
	return cost[end]
}

func neighbours(pos Pos) []Pos {
	return []Pos{{pos.x, pos.y - 1}, {pos.x + 1, pos.y}, {pos.x - 1, pos.y}, {pos.x, pos.y + 1}}
}

func big(data [][]int, scale int) [][]int {
	board := make([][]int, 0, len(data)*scale)
	k := 0
	for y := 0; y < len(data)*scale; y++ {
		board = append(board, make([]int, 0, len(data[0])*scale))
		for x := 0; x < len(data[0])*scale; x++ {
			val := data[y%len(data)][x%len(data[0])]
			val += y/len(data) + x/len(data[0])
			for val > 9 {
				val -= 9
			}
			board[k] = append(board[k], val)
		}
		k++
	}
	return board

}

type item struct {
	pos Pos
	val int
}
type priorityQueue []item

func (q priorityQueue) Len() int { return len(q) }
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q priorityQueue) Less(i, j int) bool {
	return q[i].val < q[j].val
}
func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(item))
}
func (q *priorityQueue) Pop() interface{} {
	n := len(*q)
	it := (*q)[n-1]
	*q = (*q)[:n-1]
	return it
}
