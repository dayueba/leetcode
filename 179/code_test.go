package main

import (
	"fmt"
	"sort"
)

func compare(a, b int) bool {
	for a > 0 && b > 0 {

	}
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(i, j)
	})

	fmt.Println(nums)
}

func main() {

}
