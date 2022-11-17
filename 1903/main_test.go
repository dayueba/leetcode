package main

import (
	"fmt"
	"testing"
)

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		n := int(num[i] - '0')
		if n%2 == 1 {
			return num[0 : i+1]
		}
	}

	return ""
}

func TestLargestOddNumber(t *testing.T) {
	n := "35427"
	//fmt.Println(n[0:5])
	fmt.Println(largestOddNumber(n))
}
