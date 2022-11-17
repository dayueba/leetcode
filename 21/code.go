package main

import (
	"fmt"
	"leetcode/utils"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	fakeHead := &utils.ListNode{}
	c := fakeHead

	for list1 != nil && list2 != nil {
		// fmt.Println(list1.Val, list2.Val)
		if list1.Val <= list2.Val {
			c.Next = list1
			list1 = list1.Next
			c = c.Next
		} else {
			c.Next = list2
			list2 = list2.Next
			c = c.Next
		}
	}
	utils.PrintList(fakeHead.Next)
	fmt.Println(list1, list2)

	if list1 != nil {
		c.Next = list1
	}
	if list2 != nil {
		fmt.Println(111)
		c.Next = list2
	}
	utils.PrintList(fakeHead.Next)

	return fakeHead.Next
}

func main() {
	mergeTwoLists(utils.MakeList([]int{1, 2, 4}), utils.MakeList([]int{1, 3, 4}))
}
