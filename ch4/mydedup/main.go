package main

import (
	"fmt"
)

func dedup(strings []string) []string {
	result := strings[:0]
	for _, str := range strings {
		if len(result) != 0 && result[len(result)-1] == str {
			continue
		}
		result = append(result, str)
	}
	return result
}

func main() {
	fmt.Println(dedup(nil))
	fmt.Println(dedup([]string{}))
	fmt.Println(dedup([]string{"hello", "hello", "world", "big", "big", "world"}))
	fmt.Println(dedup([]string{"a", "b", "c", "d", "e", "f"}))
	fmt.Println(dedup([]string{"абвгд", "абвгд", "абвгд", "абвг", "абвгд"}))
}
