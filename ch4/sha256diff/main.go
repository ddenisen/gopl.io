package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// shaDiff takes the corresponding sha256 hashes of the two inputs
// and returns the number of bits that are different between them
func shaDiff(data1, data2 []byte) int {
	hash1 := sha256.Sum256(data1)
	hash2 := sha256.Sum256(data2)

	diffCount := 0
	for i := range hash1 {
		diffCount += int(pc[hash1[i]] ^ pc[hash2[i]])
	}
	return diffCount
}

func main() {
	fmt.Println(shaDiff([]byte("abc"), []byte("abc")))
	fmt.Println(shaDiff([]byte("x"), []byte("X")))
}
