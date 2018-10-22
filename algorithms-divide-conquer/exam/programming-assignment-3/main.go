package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var x int
	var input []int

	for {
		_, err := fmt.Fscan(r, &x)
		if err != nil {
			break
		}
		input = append(input, x)
	}

	a := make([]int, len(input))

	copy(a, input)
	fmt.Println(quickSort(a, alwaysFirst))
	// fmt.Println(a)

	copy(a, input)
	fmt.Println(quickSort(a, alwaysLast))
	// fmt.Println(a)

	copy(a, input)
	fmt.Println(quickSort(a, medianOfThree))
	// fmt.Println(a)
}

func quickSort(a []int, pivot func([]int, int, int) int) int {
	return _quickSort(a, 0, len(a), pivot)
}

func _quickSort(a []int, l, h int, pivot func([]int, int, int) int) int {
	n := h - l

	if n <= 1 {
		return 0
	}

	p := pivot(a, l, h)
	a[l], a[p] = a[p], a[l]

	i := l + 1
	for j := l + 1; j < h; j++ {
		if a[j] < a[l] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[l], a[i-1] = a[i-1], a[l]

	return (n - 1) + _quickSort(a, l, i-1, pivot) + _quickSort(a, i, h, pivot)
}

func alwaysFirst(a []int, l, h int) int {
	return l
}

func alwaysLast(a []int, l, h int) int {
	return h - 1
}

func medianOfThree(a []int, l, h int) int {
	n := h - l
	m := l + n/2
	if n%2 == 0 {
		m--
	}

	ll := a[l]
	mm := a[m]
	hh := a[h-1]

	b := []int{ll, mm, hh}
	sort.Ints(b)

	switch b[1] {
	case ll:
		return l
	case mm:
		return m
	case hh:
		return h - 1
	}

	return -1
}
