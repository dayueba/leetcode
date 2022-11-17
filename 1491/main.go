package main

import "fmt"

func average(salary []int) float64 {
	max := 1000
	min := 1000000
	sum := 0
	for i := 0; i < len(salary); i++ {
		val := salary[i]
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
		sum += val
	}
	return float64(sum-min-max) / float64(len(salary) - 2)
}

func main() {
	arr := []int{2000,375000,349000,9000,335000,423000,428000,484000,482000,145000,328000,72000,470000,275000,55000,448000,182000,128000,475000,368000,469000,268000,265000,397000,323000,245000,173000,460000,183000,404000,123000,248000,295000,4000,27000,281000,413000,218000}
	fmt.Println(average(arr))
}
