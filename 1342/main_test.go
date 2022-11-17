package main

import (
	"testing"
)

func numberOfSteps(num int) int {
	ans := 0
	for ; num > 0; num >>= 1 {
		ans += num & 1 // 奇数额外加1
		if num > 1 {
			ans++
		}
	}
	return ans
}

func TestNumberOfSteps(t *testing.T) {
	if numberOfSteps(14) != 6 {
		t.Error()
	}
}
