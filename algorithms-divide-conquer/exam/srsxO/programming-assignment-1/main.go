/*
Package karatsuba implements the Karatsuba algorithm for my homework.
*/
package karatsuba

import (
	"bytes"
	"strconv"
)

// Karatsuba multiplies 2 large integers in the Karatsuba algorithm.
func Karatsuba(x, y string) string {
	_x := make([]int, 0)
	_y := make([]int, 0)

	for i := len(x) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(x[i]))
		_x = append(_x, n)
	}
	for i := len(y) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(y[i]))
		_y = append(_y, n)
	}

	z := mul(_x, _y)

	var buf bytes.Buffer
	for i := len(z) - 1; i >= 0; i-- {
		buf.WriteString(strconv.Itoa(z[i]))
	}
	return buf.String()
}

func mul(x, y []int) []int {
	n, xn, yn := 0, len(x), len(y)
	if xn > n {
		n = xn
	}
	if yn > n {
		n = yn
	}

	// Stop the recursion.
	if n == 1 {
		return mul1(x, y)
	}

	// Normalize n as an even number.
	n += n % 2

	a := x[n/2:]
	b := x[:n/2]
	c := y[n/2:]
	d := y[:n/2]

	ac := mul(a, c)
	bd := mul(b, d)

	// (a+b)(c+d)
	abcd := mul(add(a, b), add(c, d))
	// ad+bc = (a+b)(c+d)-bd-ac
	adbc := sub(sub(abcd, bd), ac)

	return add(
		add(shift(ac, n), bd),
		shift(adbc, n/2),
	)
}

func get(a []int, i int) int {
	if len(a) > i {
		return a[i]
	}
	return 0
}

func shift(n []int, shift int) []int {
	m := make([]int, shift+len(n))

	for i := 0; i < shift; i++ {
		m[i] = 0
	}

	copy(m[shift:], n)

	return m
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

func mul1(x, y []int) []int {
	z := get(x, 0) * get(y, 0)
	if z < 10 {
		return []int{z}
	}
	return []int{z % 10, z / 10}
}
