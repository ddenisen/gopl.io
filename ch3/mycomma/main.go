package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func comma(s string) string {
	// n tracks the number of non-fractional digits in s
	var n int
	if i := strings.IndexByte(s, '.'); i >= 0 {
		n = i
	} else {
		n = utf8.RuneCountInString(s)
	}

	// ds denotes the starting index of the first digit in s
	var ds = 0
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		ds++
	}

	var buf bytes.Buffer
	for i, r := range s {
		if ds < i && i < n && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("0"))
	fmt.Println(comma("1"))
	fmt.Println(comma("-1"))
	fmt.Println(comma("18"))
	fmt.Println(comma("-34"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("22346"))
	fmt.Println(comma("-1001"))
	fmt.Println(comma("-200000"))
	fmt.Println(comma("1000000"))
	fmt.Println(comma("0.1"))
	fmt.Println(comma("-14.65"))
	fmt.Println(comma("1000.99999"))
	fmt.Println(comma("1.5544e8"))
	fmt.Println(comma("-3455.4344433"))
}
