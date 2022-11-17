package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte, len(s))
	t2s := make(map[byte]byte, len(t))

	for idx := 0; idx < len(s); idx++ {
		x, y := s[idx], t[idx]
		if s2t[x] > 0 && s2t[x] != y {
			return false
		}
		if t2s[y] > 0 && t2s[y] != x {
			return false
		}

		s2t[x] = y
		t2s[y] = x
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add") == true)
	fmt.Println(isIsomorphic("foo", "bar") == false)
}
