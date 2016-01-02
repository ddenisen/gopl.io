package main

import (
	"fmt"
)

func countRunes(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	return count
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Count := countRunes(s1)
	s2Count := countRunes(s2)

	for r, n := range s1Count {
		if s2Count[r] != n {
			return false
		}
	}

	return true
}

func test(s1, s2 string, exp bool) {
	actual := anagram(s1, s2)
	if actual == exp {
		//fmt.Printf("anagram(\"%s\",\"%s\") correctly returned %t\n", s1, s2, actual)
	} else {
		fmt.Printf("FAIL anagram(\"%s\",\"%s\") incorrectly returned %t\n", s1, s2, actual)
	}
}

func main() {
	test("", "", true)
	test("", "  ", false)
	test("a", "b", false)
	test("abra", "bara", true)
	test("abra", "abra", true)
	test("abra", "", false)
	test("abracadabra", "candelabra", false)
	test("abracadabra", "abra", false)
	test("abracadabra", "embarcadero", false)
	test("hatmantan", "manhattan", true)
	test("квартира", "дом", false)
	test("улица", "лицау", true)
	test("улица", "улика", false)
}
