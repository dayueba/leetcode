package _832

import (
	"testing"
)

func checkIfPangram(sentence string) bool {
	m := map[int]int{}
	for _, v := range []rune(sentence) {
		m[int(v)-96] = 1
	}
	for i := 1; i < 27; i++ {
		if m[i] == 0 {
			return false
		}
	}
	return true
}

func TestCheckIfPangram(t *testing.T) {
	checkIfPangram("thequickbrownfoxjumpsoverthelazydog")
}
