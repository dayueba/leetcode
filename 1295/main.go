package main

import "fmt"

func findNumbers(nums []int) int {
	n := 0
	for _, num := range nums {
		size := 1
		for num/10 > 0 {
			size += 1
			num /= 10
		}
		if size%2 == 0 {
			n += 1
		}
	}
	return n
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
}
