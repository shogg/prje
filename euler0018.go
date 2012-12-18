package main

import (
	"strings"
	"strconv"
	"fmt"
	"container/heap"
)

var (
	graphString =
`75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
)

type node struct {
	val int
	prev int
	mean float64
}

type graph []node
type path []int

type pqueue struct {
	g graph
	nodes []int
	indexes []int
}

func main() {
	g := parseGraph(graphString)
	g[3].val = 666
	g.dijkstra()
	fmt.Print(g)
}

func (g graph) dijkstra() {
	var pq *pqueue = newPqueue(g)
	heap.Init(pq)

	g[0].mean = float64(g[0].val)
	for i := 0; i < len(g); i++ {
		heap.Push(pq, i)
	}

	fmt.Print(g)

	for pq.Len() > 0 {
		n := heap.Pop(pq).(int)
		println("pop", n)
		for _, m := range g.neighbors(n) {
			if pq.Indexof(m) > 0 {
				g.update(m, n, pq)
			}
		}
	}
}

func newPqueue(g graph) *pqueue {
	n := len(g)
	var pq *pqueue = new(pqueue)
	pq.g = g
	pq.nodes = make([]int, 0, n)
	pq.indexes = make([]int, n)
	return pq
}

func (g graph) update(node, prev int, pq *pqueue) {

	mean := float64(g[node].val)
	count := 1
	for n := prev; n >= 0; n = g[n].prev {
		mean += float64(g[n].val)
		count++
	}

	mean /= float64(count)

	println("update", node)

	if(mean > g[node].mean) {
		g[node].mean = mean
		g[node].prev = prev
	}
	fmt.Println(g)

	i := pq.Indexof(node)
	println(i)

	removed := heap.Remove(pq, i).(int)
	if(removed != node) {
		panic(fmt.Sprint("panicker ", removed, " != ", node))
	}

	heap.Push(pq, node)
}

func (g graph) neighbors(n int) []int {

	row := 0
	for i := 0; 2*i + 1 < n; i = 2*i + 1 {
		row++
	}

	var m []int
	if row + 1 >= 15 { return nil }

	m = []int { n + row + 1, n + row + 2 }

	println(n, "->", m[0], m[1])
	return m
}

func (pq *pqueue) Push(x interface{}) {
	n := pq.Len()
	pq.nodes = append(pq.nodes, x.(int))
	pq.indexes[x.(int)] = n
}

func (pq *pqueue) Pop() interface{} {
	n := pq.Len()
	x := pq.nodes[n-1]
	pq.nodes = pq.nodes[0:n-1]
	pq.indexes[x] = -1
	return x
}

func (pq *pqueue) Len() int { return len(pq.nodes) }

func (pq *pqueue) Less(i, j int) bool {
	return pq.g[i].mean > pq.g[j].mean
}
func (pq *pqueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}

func (pq *pqueue) Indexof(x int) int {
	return pq.indexes[x]
}

func parseGraph(s string) graph {
	var g graph
	rows := strings.Split(s, "\n")
	for _, row := range rows {
		values := strings.Split(row, " ")
		for _, val := range values {
			num, _ := strconv.Atoi(val)
			g = append(g, node { num, -1, 0.0 })
		}
	}
	return g
}

func (g graph) String() string {
	s := ""
	i := 0
	for row := 0; row < 15; row++ {
		for col := 0; col <= row; col++ {
			s += fmt.Sprint(g[i]) + " "
			i++
		}
		s += "\n"
	}
	return s
}
