package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canBeEqual(target []int, arr []int) bool {
	m := map[int]int{}
	for _, v := range target {
		m[v]++
	}
	for _, v := range arr {
		if m[v] <= 0 {
			return false
		}
		m[v]--
	}
	return true
}

func TestSomething(t *testing.T) {
	assert.Equal(t, canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}), true, "they should be equal")
	assert.Equal(t, canBeEqual([]int{7}, []int{7}), true, "they should not be equal")
	assert.Equal(t, canBeEqual([]int{3, 7, 9}, []int{3, 7, 11}), false, "they should be equal")
}
