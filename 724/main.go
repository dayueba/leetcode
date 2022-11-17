package main

import "fmt"

func pivotIndex(nums []int) int {
	nums1, nums2 := make([]int, len(nums)+1), make([]int, len(nums)+1)
	nums1[0] = 0
	nums1[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		nums1[i] = nums[i-1] + nums1[i-1]
	}
	sum := nums1[len(nums)]
	nums2[0] = nums1[len(nums)]
	for i := 1; i < len(nums); i++ {
		sum -= nums[i-1]
		nums2[i] = sum
	}

	for i := 0; i < len(nums); i++ {
		if nums1[i] == nums2[i+1] {
			return i
		}
	}
	return -1
}

func pivotIndex2(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}

// [1, 7, 3, 6, 5, 6]
// 0, 1, 8, 11, 17, 22, 28]
// 28 27 20 17  11  5   0

// [2, 1, -1]
// [0. 2. 3, 2]
// [2, 0. -1. 0]
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	fmt.Println(pivotIndex([]int{2, 1, 3}))
}
