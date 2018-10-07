package karatsuba

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpad(t *testing.T) {
	assert.Equal(t, []int{}, unpad([]int{0, 0, 0, 0, 0, 0}))
	assert.Equal(t, []int{0, 0, 0, 1}, unpad([]int{0, 0, 0, 1, 0, 0}))
}

func TestShift(t *testing.T) {
	assert.Equal(t, []int{0, 0, 1}, shift([]int{1}, 2))
	assert.Equal(t, []int{0, 0, 1, 2, 3}, shift([]int{1, 2, 3}, 2))
}

func TestAdd(t *testing.T) {
	assert.Equal(t, []int{5, 7, 9}, add([]int{1, 2, 3}, []int{4, 5, 6}))
	assert.Equal(t, []int{9, 9, 9}, add([]int{1}, []int{8, 9, 9}))
	assert.Equal(t, []int{0, 0, 0, 1}, add([]int{1}, []int{9, 9, 9}))
}

func TestSub(t *testing.T) {
	assert.Equal(t, []int{3, 3, 3}, sub([]int{4, 5, 6}, []int{1, 2, 3}))
	assert.Equal(t, []int{2, 2, 2}, sub([]int{1, 2, 3}, []int{9, 9}))
	assert.Equal(t, []int{1}, sub([]int{0, 1}, []int{9}))
}

func TestMul(t *testing.T) {
	assert.Equal(t, []int{6, 5}, mul([]int{8}, []int{7}))
	assert.Equal(t, []int{0, 0, 2}, mul([]int{0, 1}, []int{0, 2}))
	assert.Equal(t, []int{1, 2, 2, 1}, mul([]int{1, 1}, []int{1, 1, 1}))
	assert.Equal(t, []int{4, 6, 1, 6}, mul([]int{6, 4}, []int{4, 3, 1}))
	assert.Equal(t, []int{2, 5, 6, 6, 0, 0, 7}, mul([]int{4, 3, 2, 1}, []int{8, 7, 6, 5}))
}

func TestKaratsuba(t *testing.T) {
	assert.Equal(t, "7006652", Karatsuba("1234", "5678"))
}
