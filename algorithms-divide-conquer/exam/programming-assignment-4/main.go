package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	graph := make(map[int][]int)

	s := bufio.NewScanner(os.Stdin)
	var line string

	for s.Scan() {
		line = s.Text()
		fields := strings.Fields(line)

		v, _ := strconv.Atoi(fields[0])
		graph[v] = []int{}

		for _, f := range fields[1:] {
			x, _ := strconv.Atoi(f)
			graph[v] = append(graph[v], x)
		}
	}

	fmt.Println(karger(graph))
}

func karger(graph map[int][]int) int {
	n := len(graph)
	minCut := -1

	for i := 0; i < n; i++ {
		g := copyGraph(graph)

		for len(g) > 2 {
			v, u := randomEdge(g)
			contractEdge(g, v, u)
		}

		for _, l := range g {
			if minCut == -1 || minCut > len(l) {
				minCut = len(l)
			}
			break
		}
	}

	return minCut
}

func copyGraph(graph map[int][]int) map[int][]int {
	copied := make(map[int][]int)

	for v, l := range graph {
		copied[v] = make([]int, len(l))
		copy(copied[v], l)
	}

	return copied
}

func contractEdge(graph map[int][]int, v, u int) {
	// merge u into v
	graph[v] = append(graph[v], graph[u]...)
	delete(graph, u)

	// redirect u to v
	for _, l := range graph {
		for i, x := range l {
			if x == u {
				l[i] = v
			}
		}
	}

	// remove loops
	l := []int{}
	for _, x := range graph[v] {
		if x != v {
			l = append(l, x)
		}
	}
	graph[v] = l
}

func randomEdge(graph map[int][]int) (int, int) {
	vs := make([]int, len(graph))

	i := 0
	for v := range graph {
		vs[i] = v
		i++
	}

	v := vs[rand.Intn(len(vs))]
	u := graph[v][rand.Intn(len(graph[v]))]

	return v, u
}
