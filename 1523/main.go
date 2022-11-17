package main

import "fmt"

func countOdds(low int, high int) int {
	count := 0
	for ; low <= high; low++ {
		if low % 2 != 0 {
			count += 1
		}
	}

	return count
}

func countOdds2(low int, high int) int {
	count := 0
	if low % 2 != 0 {
		count += 1
	}
	if high % 2 != 0 {
		count += 1
	}

	return count + (high - low) / 2
}

func countOdds3(low int, high int) int {
	return pre(high) - pre(low-1)
}

func pre(x int) int {
	return(x + 1) >> 1
}

func main() {
	//fmt.Println(countOdds(3,7))
	//fmt.Println(countOdds(8,10))
	//fmt.Println(countOdds2(3,7))
	//fmt.Println(countOdds2(8,10))
	fmt.Println(countOdds3(3,7))
	fmt.Println(countOdds3(8,10))
}
