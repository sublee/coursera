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
	_x := atol(x)
	_y := atol(y)

	_z := mul(_x, _y)
	z := ltoa(_z)

	return z
}

// atol converts a number string to a reversed slice of digits.
//
//   atol("123") -> []int{3, 2, 1}
//
func atol(a string) []int {
	l := make([]int, 0)

	for i := len(a) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(a[i]))
		l = append(l, n)
	}

	return l
}

// ltoa converts a reversed slice of digits to a number string.
//
//   ltoa([]int{3, 2, 1}) -> "123"
//
func ltoa(l []int) string {
	var buf bytes.Buffer

	for i := len(l) - 1; i >= 0; i-- {
		buf.WriteString(strconv.Itoa(l[i]))
	}

	return buf.String()
}

// min returns the smaller number.
//
//   min(1, 2) -> 1
//
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// get returns the number at the offset or 0.
//
//   get([]int{1}, 0) -> 1
//   get([]int{1}, 1) -> 0
//
func get(l []int, i int) int {
	if len(l) > i {
		return l[i]
	}
	return 0
}

// shift prepend zeros into a reversed slice of digits.
//
//   shift([]int{1}, 2) -> []int{0, 0, 2}
//
func shift(l []int, shift int) []int {
	m := make([]int, shift+len(l))

	for i := 0; i < shift; i++ {
		m[i] = 0
	}

	copy(m[shift:], l)

	return m
}

// unpad discards trailing zeros.
//
//   unpad([]int{0, 0, 1, 0, 0}) -> []int{0, 0, 1}
//
func unpad(l []int) []int {
	for i := len(l) - 1; i >= 0; i-- {
		if l[i] != 0 {
			return l[:i+1]
		}
	}
	return l[:0]
}

// add adds two reversed slices of digits.
//
//   sub([]int{2, 4}, []int{4, 2}) -> []int{6, 6}
//
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

// sub subtracts two reversed slices of digits.
//
//   sub([]int{2, 4}, []int{4, 2}) -> []int{8, 1}
//
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

// mul multiplies two reversed slices of digits in the Karatsuba algorithm.
//
//   mul1([]int{2, 4}, []int{4, 2}) -> []int{8, 0, 0, 1}
//
func mul(x, y []int) []int {
	n, xn, yn := 0, len(x), len(y)
	if xn > n {
		n = xn
	}
	if yn > n {
		n = yn
	}

	// Stop the recursion.
	switch n {
	case 0:
		return []int{0}
	case 1:
		z := get(x, 0) * get(y, 0)
		return unpad([]int{z % 10, z / 10})
	}

	// Normalize n as an even number.
	n += n % 2

	a := x[min(len(x), n/2):]
	b := x[:n/2]
	c := y[min(len(y), n/2):]
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
