package main

import "fmt"

func subtractProductAndSum(n int) int {
	n1, n2 := 0, 1
	for ; n > 0; n /= 10 {
		val := n % 10
		n1 += val
		n2 *= val
	}

	return n2 - n1
}

func main() {
	fmt.Println(subtractProductAndSum(234))
}
