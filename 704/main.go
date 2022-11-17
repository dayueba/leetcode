package main

import "fmt"

func search(nums []int, target int) int {
	// []
	i, j := 0, len(nums)-1

	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			j = mid - 1

		}
		if nums[mid] < target {
			i = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9) == 4)
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2) == -1)

}
