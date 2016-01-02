package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func comma(s string) string {
	n := utf8.RuneCountInString(s)
	var buf bytes.Buffer
	for i, r := range s {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("22346"))
	fmt.Println(comma("1000000"))
}
