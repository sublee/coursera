package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var x int
	var ints []int

	for {
		_, err := fmt.Fscan(r, &x)
		if err != nil {
			break
		}
		ints = append(ints, x)
	}

	fmt.Println(solve(ints))
}

func solve(ints []int) int {
	return sortAndCountInv(ints, 0, len(ints))
}

// sortAndCountInv is a Merge Sort implementation but it counts inversions
// together.
func sortAndCountInv(ints []int, l, h int) int {
	n := h - l

	if n == 1 {
		return 0
	}

	m := (l + h) / 2

	a := sortAndCountInv(ints, l, m)
	b := sortAndCountInv(ints, m, h)
	c := 0

	out := make([]int, n)
	i, j := l, m

	for k := 0; k < n; k++ {
		if i == m {
			copy(out[k:n], ints[j:h])
			break
		}
		if j == h {
			copy(out[k:n], ints[i:m])
			break
		}

		if ints[j] < ints[i] {
			out[k] = ints[j]
			j++

			// Count inversions.
			c += m - i
			// When choose 2 against 3 in [1 3 5] [2 4 6], there are 2
			// inversions: (3 2) and (5 2). That's why here c should be
			// increased by m-i instead of just 1.

		} else {
			out[k] = ints[i]
			i++
		}
	}

	copy(ints[l:h], out)

	return a + b + c
}
