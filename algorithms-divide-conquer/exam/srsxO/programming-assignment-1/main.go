package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Karatsuba("123", "456"))
}

func Karatsuba(x, y string) string {
	x_ := make([]int, 0)
	y_ := make([]int, 0)

	for i := len(x); i >= 0; i-- {
		n, _ := strconv.Atoi(string(x[i]))
		x_ = append(x_, n)
	}
	for i := len(y); i >= 0; i-- {
		n, _ := strconv.Atoi(string(y[i]))
		y_ = append(y_, n)
	}

	z := karatsuba(x_, y_)

	var buf bytes.Buffer
	for i := len(z); i >= 0; i-- {
		buf.WriteString(strconv.Itoa(z[i]))
	}
	return buf.String()
}

func karatsuba(x, y []int) []int {
	n := len(x)
	// if len(y) > n {
	// 	n = len(y)
	// }

	a := x[:n/2]
	b := x[n/2:]

	c := y[:n/2]
	d := y[n/2:]

	ac := karatsuba(a, c)
	bd := karatsuba(b, d)
	abcd := karatsuba(add(a, b), add(c, d))
	sub(sub(abcd, bd), ac)

	return nil
}

func get(a []int, i int) int {
	if len(a) > i {
		return a[i]
	}
	return 0
}

func unpad(n []int) []int {
	for i := len(n) - 1; i >= 0; i-- {
		if n[i] != 0 {
			return n[:i+1]
		}
	}
	return n[:0]
}

func add(x, y []int) []int {
	z := make([]int, 0)
	var sum int

	for i := 0; i < len(x) || i < len(y); i++ {
		sum = get(z, i) + get(x, i) + get(y, i)

		if len(z) == i {
			z = append(z, 0)
		}

		z[i] = sum % 10
		if sum >= 10 {
			z = append(z, sum/10)
		}
	}

	return unpad(z)
}

func sub(x, y []int) []int {
	z := make([]int, 0)
	var diff int

	for i := 0; i < len(x) || i < len(y); i++ {
		diff = get(x, i) - get(y, i) - get(z, i)

		if len(z) == i {
			z = append(z, 0)
		}

		if diff < 0 {
			z[i] = 10 + diff
			z = append(z, 1)
		} else {
			z[i] = diff
		}
	}

	return unpad(z)
}
