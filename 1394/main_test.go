package main

import (
	"fmt"
	"testing"
)

func findLucky(arr []int) int {
	m := map[int]int{}

	for _, v := range arr {
		m[v]++
	}

	max := -1
	for key, val := range m {
		if key == val && key > max {
			max = key
		}
	}

	return max
}

func TestFindLucky(t *testing.T) {
	fmt.Println(findLucky())
}
