package main

import (
	"fmt"
)

func rotate(slice []int, offset int) []int {
	offset = offset % len(slice)
	return append(slice[offset:], slice[:offset]...)
}

func test(slice []int, offset int) {
	fmt.Printf("rotate(%v, %d) => %v\n", slice, offset, rotate(slice, offset))
}

func main() {
	test([]int{42}, 0)
	test([]int{42}, 4)
	test([]int{1, 2}, 1)
	test([]int{1, 2}, 2)
	test([]int{1, 2}, 3)
	test([]int{0, 1, 2, 3, 4, 5}, 0)
	test([]int{0, 1, 2, 3, 4, 5}, 1)
	test([]int{0, 1, 2, 3, 4, 5}, 4)
	test([]int{0, 1, 2, 3, 4, 5}, 6)
	test([]int{0, 1, 2, 3, 4, 5}, 7)
}
