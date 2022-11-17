package main

import (
	"testing"
)

func maxAscendingSum(nums []int) int {
	max := nums[0]
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmp += nums[i]
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = nums[i]
		}
	}
	if tmp > max {
		max = tmp
	}
	return max
}

func TestMaxAscendingSum(t *testing.T) {
	if maxAscendingSum([]int{3, 6, 10, 1, 8, 9, 9, 8, 9}) != 19 {
		t.Error()
	}
}
