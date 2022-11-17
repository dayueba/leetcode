package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if n > 0 && m > 0 && nums1[m-1] >= nums2[n-1] {
			m--
			nums1[i] = nums1[m]
		} else if n > 0 && m > 0 && nums1[m-1] <= nums2[n-1] {
			n--
			nums1[i] = nums2[n]
		} else if m > 0 {
			m--
			nums1[i] = nums1[m]
		} else if n > 0 {
			n--
			nums1[i] = nums2[n]
		}
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if n == 0 {
			m--
			nums1[i] = nums1[m]
		} else if m == 0 {
			n--
			nums1[i] = nums2[n]
		} else if nums1[m-1] >= nums2[n-1] {
			m--
			nums1[i] = nums1[m]
		} else if nums1[m-1] <= nums2[n-1] {
			n--
			nums1[i] = nums2[n]
		}
	}

	fmt.Println(nums1)
}

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)

	merge2([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	merge2([]int{1}, 1, []int{}, 0)
	merge2([]int{0}, 0, []int{1}, 1)
}
