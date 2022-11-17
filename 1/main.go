package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 10000)
	for i, v := range nums {
		if val, ok := m[target - v]; ok {
			return []int{i, val}
		}
		m[v] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
}


///// 7

////  {2,0}