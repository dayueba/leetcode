package main

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum1(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for a := 0; a < len(nums); a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		c := len(nums) - 1
		for b := a + 1; b < len(nums); b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for b < c && nums[a]+nums[b]+nums[c] > 0 {
				c--
			}

			if b == c {
				break
			}

			if nums[a]+nums[b]+nums[c] == 0 {
				res = append(res, []int{nums[a], nums[b], nums[c]})
			}

		}
	}

	return res
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4}))
}
