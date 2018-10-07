package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpad(t *testing.T) {
	assert.Equal(t, []int{}, unpad([]int{0, 0, 0, 0, 0, 0}))
	assert.Equal(t, []int{0, 0, 0, 1}, unpad([]int{0, 0, 0, 1, 0, 0}))
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
