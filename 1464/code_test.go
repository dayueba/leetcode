package code

import "testing"

func maxProduct(nums []int) int {
	// m2 > m1 > other

	var m1, m2 int
	if nums[0] > nums[1] {
		m2 = nums[0]
		m1 = nums[1]
	} else {
		m2 = nums[1]
		m1 = nums[2]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > m2 {
			m1 = m2
			m2 = nums[i]
		} else if nums[i] > m1 {
			m1 = nums[i]
		}
	}

	return (m1 - 1) * (m2 - 1)
}

func TestMaxProduct(t *testing.T) {

}
