package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(arr []int) *ListNode {
	fakeHead := &ListNode{}
	c := fakeHead
	for _, val := range arr {
		c.Next = &ListNode{Val: val}
		c = c.Next
	}

	return fakeHead.Next
}

func PrintList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
