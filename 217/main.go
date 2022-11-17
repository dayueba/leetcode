package main

import "fmt"

type empty struct {
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, val := range nums {
		if _, ok := m[val]; ok {
			return true
		}
		m[val] = empty{}
	}
	return false
}

func main() {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 20}
	fmt.Println(containsDuplicate(arr))
}
