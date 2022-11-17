package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchInsert(nums []int, target int) int {
	// []
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	return l
}

func TestSomething(t *testing.T) {
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 5), 2, "they should be equal")
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 2), 1, "they should not be equal")
	assert.Equal(t, searchInsert([]int{1, 3, 5, 6}, 7), 4, "they should not be equal")

}
