package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			res += int(math.Pow(2, float64(len(arr)-i-1)))
		}
	}
	return res
}

func main() {
	fmt.Println(math.Pow(2, 14))
	fmt.Println(getDecimalValue(nil)) // 18880
}
